package log

var w *AsyncFileWriter

// Init can initialize default logger
func init() {
	w = NewAsyncFileWriter("output/bilicenter.log")
	if err := w.initLogFile(); err != nil {
		panic(err)
	}
}
