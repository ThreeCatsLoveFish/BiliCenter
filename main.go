package main

import (
	"subcenter/application"
	"subcenter/application/awpush"
)

func main() {
	application.GlobalTaskCenter.Run()

	client := awpush.NewAWPushClient()
	client.Serve()
}
