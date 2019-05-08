package slog

import (
	"testing"
)

func TestSlog(t *testing.T) {

	slog, err := DefaultNew(func() SLogConfig {
		cfg := DefaultSLogConfig()
		cfg.SplitType = SPLIT_TYPE_TIME_CYCLE
		cfg.Condition = 1
		return cfg
	})
	if err != nil {
		panic(slog)
	}

	slog.Info("test info", "aaaaa")
	slog.Debug("test debug", "debug")
	slog.Warn("test warn", "warn")
	slog.Error("test error ", "error")
	slog.Fatal("test fatal", "fatal")
}
