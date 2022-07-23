package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"subcenter/infra/dto"

	"github.com/sethvargo/go-retry"
)

// SyncAction implement IConcurrency, support synchronous actions
type SyncAction struct{}

func (a *SyncAction) Exec(user Account, job *sync.WaitGroup, child IExec) []dto.MedalInfo {
	fail := make([]dto.MedalInfo, 0, len(user.medalsLow))
	for _, medal := range user.remainMedals {
		retryTime := 1
		backOff := retry.NewFibonacci(time.Duration(retryTime) * time.Second)
		backOff = retry.WithMaxRetries(10, backOff)
		if err := retry.Do(context.Background(), backOff, func(ctx context.Context) error {
			if ok := child.Do(user, medal); !ok {
				return retry.RetryableError(errors.New("action fail"))
			}
			return nil
		}); err != nil {
			fail = append(fail, medal)
		}
	}
	child.Finish(user, fail)
	job.Done()
	return fail
}

// AsyncAction implement IConcurrency, support asynchronous actions
type AsyncAction struct{}

func (a *AsyncAction) Exec(user Account, job *sync.WaitGroup, child IExec) []dto.MedalInfo {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	fail := make([]dto.MedalInfo, 0, len(user.medalsLow))
	for _, medal := range user.remainMedals {
		wg.Add(1)
		retryTime := 1
		backOff := retry.NewFibonacci(time.Duration(retryTime) * time.Second)
		backOff = retry.WithMaxRetries(10, backOff)
		go func(medal dto.MedalInfo) {
			if err := retry.Do(context.Background(), backOff, func(ctx context.Context) error {
				if ok := child.Do(user, medal); !ok {
					return retry.RetryableError(errors.New("action fail"))
				}
				return nil
			}); err != nil {
				mu.Lock()
				fail = append(fail, medal)
				mu.Unlock()
			}
			wg.Done()
		}(medal)
	}
	wg.Wait()
	child.Finish(user, fail)
	job.Done()
	return fail
}
