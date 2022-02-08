package main

import (
	"subcenter/service/task"
	"time"
)

func main() {
	makers := []task.TaskMaker{
		task.NewTaskMaker(time.Hour * 6, 0, 0),
	}
	taskCenter := task.NewTaskCenter(makers, 5)
	taskCenter.Run()
}
