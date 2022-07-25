package main

import (
	"os"
	"subcenter/application/api"
	"subcenter/application/awpush"
	"subcenter/application/passport"
	"subcenter/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "login":
			passport.LoginBili()
		}
		return
	}

	// Initialize the task center
	taskCenter := domain.NewTaskCenter()
	go taskCenter.Run()

	// Initialize awpush client
	client := awpush.NewAWPushClient()
	go client.Run()

	// Initialize server
	router := gin.Default()
	api.LoadApi(router)
	router.Run(":8000")
}
