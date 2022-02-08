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
	taskType  TaskType
	startTime time.Time
	pull.Pull
	push.Push
}

// NewTask create a new task
func NewTask(taskType TaskType, pullId, pushId int64) Task {
	return Task{
		taskType:  taskType,
		startTime: time.Now(),
		Pull:      pull.NewPull(pushId),
		Push:      push.NewPush(pushId),
	}
}

func (task Task) Execute() error {
	title, content, err := task.Obtain()
	if err != nil {
		return err
	}
	return task.Submit(title, content)
}

type TaskCenter struct {
	wait chan Task
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
