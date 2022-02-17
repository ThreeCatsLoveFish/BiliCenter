package task

import (
	"subcenter/service/pull"
	"subcenter/service/push"
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
		// TODO: add log here
		return
	}
	for _, data := range dataList {
		task.Submit(data)
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

type taskConfig struct {
	Init   bool   `config:"init"`
	Pull   string `config:"pull"`
	Push   string `config:"push"`
	Period int64  `config:"period"`
}

func getTaskMaker() []TaskMaker {
	taskConf := config.NewWithOptions("task", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
		opt.ParseEnv = true
	})
	taskConf.AddDriver(toml.Driver)
	err := taskConf.LoadFiles("config/task.toml")
	if err != nil {
		panic(err)
	}

	// Load config file
	size := taskConf.Get("global.size").(int64)
	conf := make([]taskConfig, size)
	taskConf.BindStruct("tasks", &conf)
	makers := make([]TaskMaker, 0, size)
	for _, c := range conf {
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
	wait := make(chan Task, len(makers))
	return TaskCenter{makers, wait}
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
