package slog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var logger *LoggerS

func GetLogger() *LoggerS {
	return logger
}

type LoggerS struct {
	*log.Logger
	cfg            *SLogConfig //save setting
	skip           int
	debug          bool
	split_type     uint8 //will determine LoggerS how to work
	count          int   //number of split
	max_size       int64 //the number is KB, if max size > log file size then segmentation the log file
	intervals_time int64 //segmentation log file cycle
	btime          int64 //begin time
	size           int64 //the number is byte
	writer         Writer
	mu             *sync.RWMutex
}

func (l *LoggerS) GetSkip() int {
	return l.skip + 1
}

func (l *LoggerS) SetSkip(skip int) {
	l.skip = skip
}

func (l *LoggerS) Copy() LoggerS {
	return *l
}

func (l *LoggerS) GetWriter() io.Writer {
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

func Println(i ...interface{}) {
	logger.Println(i...)
}

func Printf(format string, i ...interface{}) {
	logger.Printf(format, i...)
}

func Debug(i ...interface{}) {
	logger.Debug(i...)
}

func Debugf(format string, i ...interface{}) {
	logger.Debugf(format, i)
}

func Info(i ...interface{}) {
	logger.Info(i...)
}

func Infof(format string, i ...interface{}) {
	logger.Infof(format, i...)
}

func Error(i ...interface{}) {
	logger.Error(i...)
}

func Errorf(format string, i ...interface{}) {
	logger.Errorf(format, i...)
}

func Warn(i ...interface{}) {
	logger.Warn(i...)
}

func Warnf(format string, i ...interface{}) {
	logger.Warnf(format, i...)
}

func Fatal(i ...interface{}) {
	logger.Fatal(i...)
}

func Fatalf(format string, i ...interface{}) {
	logger.Fatalf(format, i...)
}

func Panic(i ...interface{}) {
	logger.Panic(i...)
}

func Panicf(format string, i ...interface{}) {
	logger.Panicf(format, i...)
}

func (l *LoggerS) Println(i ...interface{}) {
	err := l.Output(1<<8-1, l.skip, fmt.Sprintln("[Println]", i))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Printf(format string, i ...interface{}) {
	err := l.Output(1<<8-1, l.skip, fmt.Sprintf("[Println] "+format, i...))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Debug(i ...interface{}) {
	if l.debug {
		err := l.Output(LOG_LEVEL_DEBUG, l.skip, fmt.Sprintln("[Debug]", i))
		if err != nil {
			fmt.Println("slog output error :", err)
		}
	}
}

func (l *LoggerS) Debugf(format string, i ...interface{}) {
	if l.debug {
		err := l.Output(LOG_LEVEL_DEBUG, l.skip, fmt.Sprintf("[Debug] "+format, i...))
		if err != nil {
			fmt.Println("slog output error :", err)
		}
	}
}

func (l *LoggerS) Info(i ...interface{}) {
	err := l.Output(LOG_LEVEL_INFO, l.skip, fmt.Sprintln("[Info]", i))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Infof(format string, i ...interface{}) {
	err := l.Output(LOG_LEVEL_INFO, l.skip, fmt.Sprintf("[Info] "+format, i...))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Error(i ...interface{}) {
	err := l.Output(LOG_LEVEL_ERROR, l.skip, fmt.Sprintln("[Error]", i))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Errorf(format string, i ...interface{}) {
	err := l.Output(LOG_LEVEL_ERROR, l.skip, fmt.Sprintf("[Error] "+format, i...))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Warn(i ...interface{}) {
	err := l.Output(LOG_LEVEL_WARN, l.skip, fmt.Sprintln("[Warn]", i))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Warnf(format string, i ...interface{}) {
	err := l.Output(LOG_LEVEL_WARN, l.skip, fmt.Sprintf("[Warn] "+format, i...))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Fatal(i ...interface{}) {
	err := l.Output(LOG_LEVEL_FATAL, l.skip, fmt.Sprintln("[Fatal]", i))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Fatalf(format string, i ...interface{}) {
	err := l.Output(LOG_LEVEL_FATAL, l.skip, fmt.Sprintf("[Fatal] "+format, i...))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
}

func (l *LoggerS) Panic(i ...interface{}) {
	s := fmt.Sprint(i...)
	err := l.Output(LOG_LEVEL_FATAL, l.skip, fmt.Sprintln("[Panic]", i)+string(debug.Stack()))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
	panic(s)
}

func (l *LoggerS) Panicf(format string, i ...interface{}) {
	s := fmt.Sprint(i...)
	err := l.Output(LOG_LEVEL_FATAL, l.skip, fmt.Sprintf("[Panic] "+format, i...)+string(debug.Stack()))
	if err != nil {
		fmt.Println("slog output error :", err)
	}
	panic(s)
}
