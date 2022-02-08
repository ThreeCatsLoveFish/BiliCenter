package main

import (
	"subcenter/service/task"
	"time"
)

func main() {
	makers := []task.TaskMaker{
		task.NewTaskMaker(time.Millisecond * 10000, 0, 0),
	}
	taskCenter := task.NewTaskCenter(makers, 5)
	taskCenter.Run()
}
