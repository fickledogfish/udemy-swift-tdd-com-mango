package log

import (
	"fmt"
	"sync"
	"time"
)

type logLevel uint

const (
	LevelDebug logLevel = iota
	LevelInfo
	LevelWarning
	LevelError
)

func (l logLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"

	case LevelInfo:
		return "INFO"

	case LevelWarning:
		return "WARN"

	case LevelError:
		return "ERR"

	default:
		panic(fmt.Sprintf("Unknown log level: %d", l))
	}
}

func Debug(format string, args ...any) {
	writeLog(LevelDebug, format, args...)
}

func Info(format string, args ...any) {
	writeLog(LevelInfo, format, args...)
}

func Warn(format string, args ...any) {
	writeLog(LevelWarning, format, args...)
}

func Error(format string, args ...any) {
	writeLog(LevelError, format, args...)
}

var logWriterMutex sync.Mutex

func writeLog(level logLevel, format string, args ...any) {
	conf := getConfig()

	if level < conf.level {
		return
	}

	logWriterMutex.Lock()
	defer logWriterMutex.Unlock()

	fmt.Printf(
		"[%s] %s => %s\n",
		level.String(),
		time.Now().Format(conf.timeFormat),
		fmt.Sprintf(format, args...),
	)
}
