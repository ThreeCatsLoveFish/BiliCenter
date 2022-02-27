package main

import (
	"subcenter/application/api"
	"subcenter/application/awpush"
	"subcenter/application/frontend"
	"subcenter/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the task center
	taskCenter := domain.NewTaskCenter()
	go taskCenter.Run()

	// Initialize awpush client
	client := awpush.NewAWPushClient()
	go client.Run()

	// Initialize frontend
	router := gin.Default()
	frontend.LoadFrontend(router)
	api.LoadApi(router)
	router.Run(":8000")
}
