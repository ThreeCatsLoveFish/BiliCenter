package service

import "sync"

// Task aggregate user info and corresponding action
type Task struct {
	Account
	actions []IAction
}

func NewTask(user Account, actions []IAction) Task {
	return Task{
		Account: user,
		actions: actions,
	}
}

func (task *Task) Start() {
	wg := sync.WaitGroup{}
	for _, action := range task.actions {
		wg.Add(1)
		go action.Exec(task.Account, &wg, action)
	}
	wg.Wait()
}
