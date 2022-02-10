package main

import (
	"subcenter/service/task"
)

func main() {
	tc := task.NewTaskCenter()
	tc.Run()
}
