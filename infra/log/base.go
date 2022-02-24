package log

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

// Init can initialize default logger
func init() {
	logFile, err := os.OpenFile("output/subcenter.log",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("create log file error")
	}
	flag := log.Ldate | log.Lmicroseconds | log.Lshortfile
	logger = log.New(logFile, "", flag)
}

func Writer() io.Writer {
	return logger.Writer()
}

func Debug(format string, v ...interface{}) {
	logger.Printf("[DEBUG] "+format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Printf("[INFO] "+format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Printf("[ERROR] "+format, v...)
}
