package slog

import (
	"io"
	"os"
	"sync"
)

type Writer interface {
	io.Writer
	ReloadeFile(file *os.File)
}

type logWriter struct {
	writer io.Writer
	mu     sync.Mutex
}

func (w *logWriter) Write(b []byte) (n int, err error) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.writer != nil {
		n, err = w.writer.Write(b)
		if err != nil {
			return
		}
	}

	return
}
func (w *logWriter) ReloadeFile(file *os.File) {
	w.mu.Lock()
	c, _ := w.writer.(io.WriteCloser)
	//w.writer.Close()
	c.Close()
	w.writer = file
	w.mu.Unlock()
}
