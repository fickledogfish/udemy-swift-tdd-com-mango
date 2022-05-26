package log

import (
	"fmt"
	"time"
)

var logChan chan string

func init() {
	logChan = make(chan string, 10)

	go func(ch <-chan string) {
		for message := range ch {
			func() {
				lock, conf := getConfig()
				defer lock.Unlock()

				fmt.Fprintln(conf.writer, message)
			}()
		}
	}(logChan)
}

func Debug(format string, args ...any) {
	writeLog(logChan, LevelDebug, format, args...)
}

func Info(format string, args ...any) {
	writeLog(logChan, LevelInfo, format, args...)
}

func Warn(format string, args ...any) {
	writeLog(logChan, LevelWarning, format, args...)
}

func Error(format string, args ...any) {
	writeLog(logChan, LevelError, format, args...)
}

func writeLog(ch chan<- string, level logLevel, format string, args ...any) {
	go func() {
		lock, conf := getConfig()
		defer lock.Unlock()

		if level < conf.level {
			return
		}

		ch <- fmt.Sprintf(
			"[%s] %s => %s",
			level.String(),
			time.Now().Format(conf.timeFormat),
			fmt.Sprintf(format, args...),
		)
	}()
}
