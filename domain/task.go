package domain

import (
	"log"
	"subcenter/domain/pull"
	"subcenter/domain/push"
	"time"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

type Task struct {
	pull.Pull
	push.Push
}

func (task Task) Execute() {
	dataList, err := task.Obtain()
	if err != nil {
		log.Default().Printf("Obtain error: %v", err)
		return
	}
	for _, data := range dataList {
		if err = task.Submit(data); err != nil {
			log.Default().Printf("Submit error: %v, data: %v", err, data)
		}
	}
}

// TaskMaker can trigger new task execution
type TaskMaker struct {
	Task
	init   bool
	period time.Duration
}

// NewTaskMaker create a new task maker
func NewTaskMaker(period time.Duration, init bool, pullName, pushName string) TaskMaker {
	return TaskMaker{
		init:   init,
		period: period,
		Task: Task{
			Pull: pull.NewPull(pullName),
			Push: push.NewPush(pushName),
		},
	}
}

// createTask can create new task and send to taker
func createTask(maker TaskMaker, takers chan Task) {
	if maker.init {
		takers <- maker.Task
	}
	for {
		timer := time.NewTimer(maker.period)
		<-timer.C
		takers <- maker.Task
	}
}

type TaskConfig struct {
	Init   bool   `config:"init"`
	Pull   string `config:"pull"`
	Push   string `config:"push"`
	Period int64  `config:"period"`
}

func getTaskMaker() []TaskMaker {
	conf := config.NewWithOptions("task", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
		opt.ParseEnv = true
	})
	conf.AddDriver(toml.Driver)
	err := conf.LoadFiles("config/task.toml")
	if err != nil {
		panic(err)
	}

	// Load config file
	var taskConfig []TaskConfig
	conf.BindStruct("tasks", &taskConfig)
	makers := make([]TaskMaker, 0)
	for _, c := range taskConfig {
		makers = append(
			makers,
			NewTaskMaker(time.Duration(c.Period), c.Init, c.Pull, c.Push),
		)
	}
	return makers
}

type TaskCenter struct {
	makers []TaskMaker
	takers chan Task
}

// NewTaskCenter initialize the task center
func NewTaskCenter() TaskCenter {
	makers := getTaskMaker()
	return TaskCenter{makers, make(chan Task, len(makers))}
}

// Run will block and execute all incoming tasks
func (tc *TaskCenter) Run() {
	if len(tc.makers) <= 0 {
		return
	}
	for _, maker := range tc.makers {
		go createTask(maker, tc.takers)
	}
	for task := range tc.takers {
		go task.Execute()
	}
}
