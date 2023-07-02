package slog

import (
	"testing"
	"time"
)

func TestSlog(t *testing.T) {

	var err error
	cfg := TestSLogConfig
	cfg.SplitType = SPLIT_TYPE_TIME_CYCLE
	cfg.Condition = 1
	cfg.LogPath = "./log/"

	err = DefaultNew(cfg)
	if err != nil {
		panic(err)
	}

	for {
		Info("test info", "aaaaa")
		Infof("test info %s", "aaaaa")
		Debug("test debug", "debug")
		Debugf("test debug %s", "debug")
		Warn("test warn", "warn")
		Warnf("test warn %s", "warn")
		Error("test error ", "error")
		Errorf("test error %s", "error")
		Fatal("test fatal", "fatal")
		Fatalf("test fatal %s", "fatal")
		time.Sleep(time.Second)
	}
}
