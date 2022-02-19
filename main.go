package main

import (
	"log"
	"os"
	"subcenter/application"
	"subcenter/application/awpush"
)

// initLog initialize default logger
func initLog() {
	logFile, err := os.Create("output/log.txt")
	if err != nil {
		panic("create log file error")
	}
	log.Default().SetOutput(logFile)
}

func main() {
	initLog()

	application.GlobalTaskCenter.Run()

	client := awpush.NewAWPushClient()
	client.Serve()
}
