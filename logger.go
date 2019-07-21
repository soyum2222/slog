package slog

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var Logger *LoggerS

type LoggerS struct {
	*log.Logger
	cfg            *SLogConfig //save setting
	debug          bool
	split_type     uint8 //will determine LoggerS how to work
	count          int   //number of split
	max_size       int64 //the number is KB, if max size > log file size then segmentation the log file
	intervals_time int64 //segmentation log file cycle
	btime          int64 //begin time
	size           int64 //the number is byte
	writer         Writer
	mu             sync.RWMutex
}

func (l *LoggerS) GetWriter() Writer {
	return l.writer
}

func (l *LoggerS) SetDebug(debug bool) {
	l.debug = debug
}

func (l *LoggerS) SetMaxSize(max int64) {
	l.max_size = max
}

func (l *LoggerS) SetIntervalsTime(intervals int64) {
	l.intervals_time = intervals
}

func (l *LoggerS) SetSliceType(t uint8) {
	l.split_type = t
}

func (l *LoggerS) Output(level uint8, skip int, s string) error {

	if level < l.cfg.Level {
		return nil
	}
	flag := false
	//check split type
	switch l.split_type {
	case SPLIT_TYPE_TIME_CYCLE:
		//init btime
		if l.btime == 0 {
			l.btime = time.Now().Unix()
			break
		}

		if time.Now().Unix()-l.btime >= l.intervals_time {
			l.btime = time.Now().Unix()
			flag = true
		}

	case SPLIT_TYPE_FILE_SIZE:
		l.size = l.size + int64(len(s))
		if l.size/1024 >= l.max_size {
			flag = true
			l.size = 0
		}
	}

	if flag {
		func() {
			l.mu.Lock()
			defer l.mu.Unlock()

			filename := ""
			l.count++
			switch l.split_type {
			case SPLIT_TYPE_FILE_SIZE:
				filename = l.cfg.FileNameHandler(l.count)
			case SPLIT_TYPE_TIME_CYCLE:
				filename = l.cfg.FileNameHandler(l.count)
			}

			file, err := os.Create(filename)
			if err != nil {
				l.Logger.Println("slog error by create new file:", err)
				return
			}
			l.writer.ReloadeFile(file)
		}()

	}

	l.mu.RLock()
	defer l.mu.RUnlock()
	err := l.Logger.Output(skip, s)

	return err

}

func (l *LoggerS) Println(i ...interface{}) error {
	return l.Output(1<<8-1, 3, fmt.Sprintln("[Println]", i))
}
func (l *LoggerS) Debug(i ...interface{}) error {
	return l.Output(LOG_LEVEL_DEBUG, 3, fmt.Sprintln("[Debug]", i))
}

func (l *LoggerS) Info(i ...interface{}) error {
	return l.Output(LOG_LEVEL_INFO, 3, fmt.Sprintln("[Info]", i))
}

func (l *LoggerS) Error(i ...interface{}) error {
	return l.Output(LOG_LEVEL_ERROR, 3, fmt.Sprintln("[Error]", i))
}

func (l *LoggerS) Warn(i ...interface{}) error {
	return l.Output(LOG_LEVEL_WARN, 3, fmt.Sprintln("[Warn]", i))
}

func (l *LoggerS) Fatal(i ...interface{}) error {
	return l.Output(LOG_LEVEL_FATAL, 3, fmt.Sprintln("[Fatal]", i))
}

func (l *LoggerS) Panic(i ...interface{}) error {
	return l.Output(LOG_LEVEL_FATAL, 3, fmt.Sprintln("[Panic]", i)+string(debug.Stack()))
}
