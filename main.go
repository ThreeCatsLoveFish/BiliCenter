package main

import (
	"subcenter/application/awpush"
	"subcenter/application/frontend"
	"subcenter/domain"
)

func main() {
	// Initialize the task center
	taskCenter := domain.NewTaskCenter()
	go taskCenter.Run()

	// Initialize awpush client
	client := awpush.NewAWPushClient()
	go client.Run()

	// Initialize frontend
	router := frontend.NewFrontend()
	router.Run(":8000")
}
