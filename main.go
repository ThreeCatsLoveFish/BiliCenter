package main

import (
	"log"
	"os"
	"subcenter/application/awpush"
	"subcenter/domain"
)

// initLog initialize default logger
func initLog() {
	logFile, err := os.OpenFile("output/subcenter.log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("create log file error")
	}
	log.Default().SetOutput(logFile)
	log.Default().SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	initLog()
	// Initialize the task center
	taskCenter := domain.NewTaskCenter()
	taskCenter.Run()

	// Initialize awpush client
	client := awpush.NewAWPushClient()
	client.Serve()
}
