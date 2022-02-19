package main

import (
	"log"
	"os"
	"subcenter/application/awpush"
	"subcenter/domain"
	"time"
)

// initLog initialize default logger
func initLog() {
	logFile, err := os.OpenFile(
		"output/subcenter.log."+time.Now().Format("2006_01_02"),
		os.O_APPEND|os.O_CREATE,
		0666,
	)
	if err != nil {
		panic("create log file error")
	}
	log.Default().SetOutput(logFile)
	log.Default().SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	initLog()

	taskCenter := domain.NewTaskCenter()
	taskCenter.Run()
	
	client := awpush.NewAWPushClient()
	client.Serve()
}
