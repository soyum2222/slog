package slog

import (
	"log"
	"os"
)

// is default new function
// writer are configured by default
func DefaultNew(f func() SLogConfig) (*Logger, error) {

	cfg := f()
	logger := new(Logger)
	logger.cfg = &cfg
	logger.SetSliceType(cfg.SplitType)

	switch cfg.SplitType {
	case SPLIT_TYPE_FILE_SIZE:
		logger.SetMaxSize(cfg.Condition)
	case SPLIT_TYPE_TIME_CYCLE:
		logger.SetIntervalsTime(cfg.Condition)
	}

	logger.SetDebug(cfg.Debug)

	writer := new(logWriter)

	if cfg.FileNameHandler == nil {
		cfg.FileNameHandler = cfg.name_handler
	}
	file, err := os.Create(cfg.FileNameHandler(0))
	if err != nil {
		return nil, err
	}

	writer.file = file
	if cfg.Debug {
		writer.stdout = os.Stdout
	}
	logger.writer = writer
	logger.Logger = log.New(logger.writer, cfg.Prefix, cfg.LogFlag)

	return logger, nil
}
