package service

import (
	"subcenter/service/pull"
	"subcenter/service/push"
	"time"
)

type TaskType int64

const (
	TimerTask TaskType = iota
	EventTask
	PriceTask
)

type TaskStatus int64

const (
	// Final status
	Waiting TaskStatus = iota

	// Status of Pull
	Pulling
	PullSuccess
	PullFail

	// Status of Push
	Pushing
	PushSuccess
	PushFail
)

type Task struct {
	id        int64
	taskType  TaskType
	status    TaskStatus
	startTime time.Time
	endTime   time.Time
	pull.Pull
	push.Push
}

// NewTask create a new task
func NewTask() {
	// TODO:
}

func (task Task) DoPull(success, fail chan Task) {
	err := task.Obtain()
	if err != nil {
		fail <- task
		return
	}
	success <- task
}

func (task Task) DoPush(success, fail chan Task) {
	err := task.Submit()
	if err != nil {
		fail <- task
		return
	}
	success <- task
}

type TaskCenter struct {
	wait chan Task
	pull chan Task
	push chan Task
}

// Run will block and execute all tasks
func (tc *TaskCenter) Run() {
	// TODO:
}
