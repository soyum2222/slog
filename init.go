package slog

import (
	"log"
	"os"
	"path"
	"sync"
)

// is default new function
// writer are configured by default
func DefaultNew(cfg SLogConfig) error {

	l := new(LoggerS)
	l.cfg = &cfg
	l.SetSliceType(cfg.SplitType)

	l.SetDebug(cfg.Debug)

	writer := new(logWriter)

	cfg.FileNameHandler = cfg.name_handler
	filename := cfg.FileNameHandler(0)

	var file *os.File
	var fileInfo os.FileInfo
	var err error
	if cfg.LogPath != "" {

		fileInfo, err = os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(path.Dir(filename), os.ModePerm)
				if err != nil {
					return err
				}
				file, err = os.Create(filename)
			} else {
				return err
			}
		} else {
			file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				return err
			}
		}
	}

	switch cfg.SplitType {
	case SPLIT_TYPE_FILE_SIZE:
		l.SetMaxSize(cfg.Condition)
		if fileInfo != nil {
			l.size = fileInfo.Size()
		}
	case SPLIT_TYPE_TIME_CYCLE:
		l.SetIntervalsTime(cfg.Condition)

	}

	if err != nil {
		return err
	}

	if file == nil {
		writer.writer = MultiWriteCloser()
	} else {
		writer.writer = MultiWriteCloser(file)

	}

	l.writer = writer
	l.Logger = log.New(l.writer, cfg.Prefix, cfg.LogFlag)

	l.skip = 4
	l.mu = &sync.RWMutex{}

	logger = l

	return nil
}
