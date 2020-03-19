package slog

import (
	"log"
	"os"
	"path"
)

// is default new function
// writer are configured by default
func DefaultNew(cfg SLogConfig) error {

	logger := new(LoggerS)
	logger.cfg = &cfg
	logger.SetSliceType(cfg.SplitType)

	logger.SetDebug(cfg.Debug)

	writer := new(logWriter)

	cfg.FileNameHandler = cfg.name_handler
	filename := cfg.FileNameHandler(0)

	var file *os.File
	var file_info os.FileInfo
	var err error
	if cfg.LogPath != "" {

		file_info, err = os.Stat(filename)
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
		logger.SetMaxSize(cfg.Condition)
		if file_info != nil {
			logger.size = file_info.Size()
		}
	case SPLIT_TYPE_TIME_CYCLE:
		logger.SetIntervalsTime(cfg.Condition)

	}

	if err != nil {
		return err
	}

	if file == nil {
		writer.writer = MultiWriteCloser()
	} else {
		writer.writer = MultiWriteCloser(file)

	}

	logger.writer = writer
	logger.Logger = log.New(logger.writer, cfg.Prefix, cfg.LogFlag)

	Logger = logger

	return nil
}
