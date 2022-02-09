package task

import (
	"subcenter/service/pull"
	"subcenter/service/push"
	"time"
)

type Task struct {
	pull.Pull
	push.Push
}

func (task Task) Execute() error {
	title, content, err := task.Obtain()
	if err != nil {
		return err
	}
	if len(title) <= 0 {
		return nil
	}
	return task.Submit(title, content)
}

// TaskMaker can trigger new task execution
type TaskMaker struct {
	// Task that needs to be executed
	Task
	// initTask means create a task when maker run
	initTask bool
	// period between two tasks
	period time.Duration
}

// NewTaskMaker create a new task maker
func NewTaskMaker(period time.Duration, initTask bool, pullId, pushId int) TaskMaker {
	return TaskMaker{
		initTask: initTask,
		period:   period,
		Task: Task{
			Pull: pull.NewPull(pushId),
			Push: push.NewPush(pushId),
		},
	}
}

type TaskCenter struct {
	makers []TaskMaker
	takers chan Task
}

// NewTaskCenter initialize the task center
func NewTaskCenter(makers []TaskMaker, size int) TaskCenter {
	wait := make(chan Task, size)
	return TaskCenter{makers, wait}
}

// Add will append a new task in wait channel
func (tc *TaskCenter) Add(task Task) {
	tc.takers <- task
}

// Run will block and execute all incoming tasks
func (tc *TaskCenter) Run() {
	for _, maker := range tc.makers {
		go func(maker TaskMaker, takers chan Task) {
			if maker.initTask {
				takers <- maker.Task
			}
			for {
				timer := time.NewTimer(maker.period)
				<-timer.C
				takers <- maker.Task
			}
		}(maker, tc.takers)
	}
	for {
		select {
		case task := <-tc.takers:
			go func() {
				task.Execute()
			}()
		}
	}
}
