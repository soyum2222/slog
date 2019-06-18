package slog

import (
	"testing"
)

func TestSlog(t *testing.T) {

	var err error
	Logger, err = DefaultNew(func() SLogConfig {
		cfg := TestSLogConfig()
		cfg.SplitType = SPLIT_TYPE_FILE_SIZE
		cfg.Condition = 1
		return cfg
	})
	if err != nil {
		panic(err)
	}

	Logger.Info("test info", "aaaaa")
	Logger.Debug("test debug", "debug")
	Logger.Warn("test warn", "warn")
	Logger.Error("test error ", "error")
	Logger.Fatal("test fatal", "fatal")
}
