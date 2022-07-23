package log

import "fmt"

func Debug(format string, v ...interface{}) {
	w.rotateFile()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.logger.Output(2, fmt.Sprintf("[DEBUG] "+format, v...))
}

func Info(format string, v ...interface{}) {
	w.rotateFile()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.logger.Output(2, fmt.Sprintf("[INFO] "+format, v...))
}

func Warn(format string, v ...interface{}) {
	w.rotateFile()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.logger.Output(2, fmt.Sprintf("[WARN] "+format, v...))
}

func Error(format string, v ...interface{}) {
	w.rotateFile()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.logger.Output(2, fmt.Sprintf("[ERROR] "+format, v...))
}

func PrintColor(format string, v ...interface{}) {
	w.rotateFile()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.logger.Printf(format, v...)
}

type AsyncLogger struct{}

func (AsyncLogger) Debug(v ...interface{}) {
	Debug("%v", v)
}

func (AsyncLogger) Info(v ...interface{}) {
	Info("%v", v)
}

func (AsyncLogger) Warn(v ...interface{}) {
	Warn("%v", v)
}

func (AsyncLogger) Error(v ...interface{}) {
	Error("%v", v)
}

func (AsyncLogger) Debugf(format string, v ...interface{}) {
	Debug(format, v...)
}

func (AsyncLogger) Infof(format string, v ...interface{}) {
	Info(format, v...)
}

func (AsyncLogger) Warnf(format string, v ...interface{}) {
	Warn(format, v...)
}

func (AsyncLogger) Errorf(format string, v ...interface{}) {
	Error(format, v...)
}

func (AsyncLogger) Sync() error { return nil }
