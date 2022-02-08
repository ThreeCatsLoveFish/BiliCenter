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

type Task struct {
	id        int64
	taskType  TaskType
	startTime time.Time
	endTime   time.Time
	pull.Pull
	push.Push
}

// NewTask create a new task
func NewTask() {
	// TODO:
}

func (task Task) Execute() error {
	title, content := task.Obtain()
	return task.Submit(title, content)
}

type TaskCenter struct {
	wait  chan Task
}

// NewTaskCenter initialize the task center
func NewTaskCenter(size int) TaskCenter {
	wait := make(chan Task, size)
	return TaskCenter{wait}
}

// Add will append a new task in wait channel
func (tc *TaskCenter) Add(task Task) {
	tc.wait <- task
}

// Run will block and execute all tasks
func (tc *TaskCenter) Run() {
	for {
		select {
		case task := <-tc.wait:
			go func() {
				task.Execute()
			}()
		}
	}
}
