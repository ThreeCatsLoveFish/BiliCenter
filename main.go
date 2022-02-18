package main

import (
	"subcenter/service/task"
	// "subcenter/service/awpush"
)

func main() {
	tc := task.NewTaskCenter()
	tc.Run()

	// client := awpush.NewAWPushClient()
	// client.Serve()
}
