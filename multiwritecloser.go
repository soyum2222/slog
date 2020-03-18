package slog

import (
	"io"
	"os"
)

type multiWriteCloser struct {
	stdout      *os.File
	writeCloser []io.WriteCloser
}

func MultiWriteCloser(closer ...io.WriteCloser) io.WriteCloser {
	allWriters := make([]io.WriteCloser, 0, len(closer))
	for _, w := range closer {
		allWriters = append(allWriters, w)

	}
	return &multiWriteCloser{writeCloser: allWriters, stdout: os.Stdout}
}

func (m *multiWriteCloser) Write(p []byte) (n int, err error) {

	for _, w := range m.writeCloser {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	_, err = m.stdout.Write(p)

	return len(p), err

}

func (m *multiWriteCloser) Close() error {

	for _, v := range m.writeCloser {
		err := v.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
