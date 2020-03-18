package slog

import (
	"testing"
	"time"
)

func TestSlog(t *testing.T) {

	var err error
	err = DefaultNew(func() SLogConfig {
		cfg := TestSLogConfig()
		cfg.SplitType = SPLIT_TYPE_TIME_CYCLE
		cfg.Condition = 1
		cfg.LogPath = "./log/"
		return cfg
	}())
	if err != nil {
		panic(err)
	}

	for {

		Logger.Info("test info", "aaaaa")
		Logger.Debug("test debug", "debug")
		Logger.Warn("test warn", "warn")
		Logger.Error("test error ", "error")
		Logger.Fatal("test fatal", "fatal")
		time.Sleep(time.Second)
	}
}
