package slog

import (
	"log"
	"strconv"
	"time"
)

type SLogConfig struct {
	Debug           bool  //if true then print all level and output on console
	Level           uint8 // debug < info < warn < Error < Fatal
	SplitType       uint8
	LogFlag         int
	Condition       int64
	LogFileName     string
	FileSuffix      string
	Prefix          string
	LogPath         string
	FileNameHandler func(i int) string //when splitting file get the new file name
}

func TestSLogConfig() SLogConfig {

	return SLogConfig{
		Debug:       true,
		Condition:   24 * 60 * 60,
		SplitType:   SPLIT_TYPE_TIME_CYCLE,
		LogFileName: "applog",
		FileSuffix:  "log",
		Prefix:      "",
		LogPath:     "./",
		LogFlag:     log.Lshortfile | log.Ltime | log.Ldate,
	}
}

func DefaultSLogConfig() SLogConfig {

	return SLogConfig{
		Level:       LOG_LEVEL_INFO,
		Debug:       false,
		Condition:   24 * 60 * 60,
		SplitType:   SPLIT_TYPE_TIME_CYCLE,
		LogFileName: "applog",
		FileSuffix:  "log",
		Prefix:      "",
		LogPath:     "./",
		LogFlag:     log.Lshortfile | log.Ltime | log.Ldate,
	}
}

func (cfg *SLogConfig) name_handler(count int) string {
	filename := cfg.LogPath + cfg.LogFileName + time.Now().Format("2006-01-02") +
		"-" + strconv.Itoa(count) + "." + cfg.FileSuffix
	return filename
}
