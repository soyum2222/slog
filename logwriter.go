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
	file   *os.File
	stdout *os.File
	mu     sync.Mutex
}

func (w *logWriter) Write(b []byte) (n int, err error) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		n, err = w.file.Write(b)
		if err != nil {
			return
		}
	}
	if w.stdout != nil {
		n, err = w.stdout.Write(b)
		if err != nil {
			return
		}
	}
	return
}
func (w *logWriter) ReloadeFile(file *os.File) {
	w.mu.Lock()
	w.file.Close()
	w.file = file
	w.mu.Unlock()
}
