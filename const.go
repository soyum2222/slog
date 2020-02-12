package slog

const (
	SPLIT_TYPE_TIME_CYCLE = iota
	SPLIT_TYPE_FILE_SIZE
	SPLIT_TYPE_TIME_DALIY
)

//debug < info < warn < Error < Fatal
const (
	LOG_LEVEL_DEBUG = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
)
