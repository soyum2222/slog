package slog

import (
	"testing"
	"time"
)

func TestSlog(t *testing.T) {

	var err error
	err = DefaultNew(func() SLogConfig {
		cfg := TestSLogConfig
		cfg.SplitType = SPLIT_TYPE_TIME_CYCLE
		cfg.Condition = 1
		cfg.LogPath = "./log/"
		return cfg
	}())
	if err != nil {
		panic(err)
	}

	for {

		Info("test info", "aaaaa")
		Debug("test debug", "debug")
		Warn("test warn", "warn")
		Error("test error ", "error")
		Fatal("test fatal", "fatal")
		time.Sleep(time.Second)
	}
}
