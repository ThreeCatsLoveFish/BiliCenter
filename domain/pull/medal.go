package pull

import (
	"subcenter/application/medalhelper/service"
	"subcenter/domain/push"
	"subcenter/infra/conf"
	"sync"
)

type MedalPull struct {
	accounts []service.Account
}

func NewMedalPull() MedalPull {
	pull := MedalPull{}
	for _, user := range conf.BiliConf.Users {
		account := service.NewAccount(user)
		pull.accounts = append(pull.accounts, account)
	}
	return pull
}

func (pull MedalPull) Obtain() ([]push.Data, error) {
	wg := sync.WaitGroup{}
	for _, account := range pull.accounts {
		wg.Add(1)
		go func(user service.Account, wg *sync.WaitGroup) {
			if status := user.Init(); status {
				user.Start(wg)
			}
		}(account, &wg)
	}
	wg.Wait()
	return nil, nil
}
