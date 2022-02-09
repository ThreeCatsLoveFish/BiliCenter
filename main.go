package main

import (
	"subcenter/service/task"
	"time"
)

func main() {
	heartbeat := task.NewTaskMaker(time.Hour*1, true, 0, 0)
	makers := []task.TaskMaker{
		heartbeat,
	}
	taskCenter := task.NewTaskCenter(makers, 5)
	taskCenter.Run()
}
