package main

import (
	"log"
	"os"
	"subcenter/application"
	"subcenter/application/awpush"
	"time"
)

// initLog initialize default logger
func initLog() {
	now := time.Now().Format("2006_01_02.15")
	logFile, err := os.Create("output/subcenter.log." + now)
	if err != nil {
		panic("create log file error")
	}
	log.Default().SetOutput(logFile)
	log.Default().SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	initLog()

	application.GlobalTaskCenter.Run()

	client := awpush.NewAWPushClient()
	client.Serve()
}
