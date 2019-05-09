# slog
a simple log with file splitting

***

# DEMO
```
	slog, err := DefaultNew(func() SLogConfig {
		cfg := TestSLogConfig()
		cfg.SplitType = SPLIT_TYPE_FILE_SIZE
		cfg.Condition = 2<<20
		return cfg
	})
	if err != nil {
		panic(err)
	}
	slog.Info("test info", "aaaaa")
	slog.Debug("test debug", "debug")
	slog.Warn("test warn", "warn")
	slog.Error("test error ", "error")
	slog.Fatal("test fatal", "fatal")
```
# OUTPUT
```
2019/05/08 18:15:08 slog_test.go:19: [Info] [test info aaaaa]
2019/05/08 18:15:08 slog_test.go:20: [Debug] [test debug debug]
2019/05/08 18:15:08 slog_test.go:21: [Warn] [test warn warn]
2019/05/08 18:15:08 slog_test.go:22: [Error] [test error  error]
2019/05/08 18:15:08 slog_test.go:23: [Fatal] [test fatal fatal]
```

### SLogconfig
```
type SLogConfig struct {
	Debug           bool  //if true then print all level and output on console
	Level           uint8 // debug < info < warn < Error < Fatal
	SplitType       uint8 // 
	LogFlag         int   // just like log packge flag
	Condition       int64 // when you use size split it is max size ,time like too
	LogFileName     string
	FileSuffix      string
	Prefix          string// just like log packge prefix
	LogPath         string
	FileNameHandler func(i int) string //when splitting file get the new file name 
}
```
***

##### if you want to rename the split file name,you need recode and assignment to FileNameHandler
##### original FileNameHandler looks like this
```
func (cfg *SLogConfig) name_handler(count int) string {
	filename := cfg.LogPath + cfg.LogFileName + time.Now().Format("2006-01-02") +
		"-" + strconv.Itoa(count) + "." + cfg.FileSuffix
	return filename
}
```
