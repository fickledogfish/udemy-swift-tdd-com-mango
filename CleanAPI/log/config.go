package log

import (
	"io"
	"os"
	"sync"
	"time"
)

type logConfig struct {
	level      logLevel
	timeFormat string
	writer     io.Writer
}

var (
	conf           logConfig
	initLogConfig  sync.Once
	getConfigMutex sync.Mutex
)

func getConfig() (*sync.Mutex, *logConfig) {
	initLogConfig.Do(func() {
		conf = logConfig{
			level:      LevelDebug,
			timeFormat: time.RFC3339,
			writer:     os.Stdout,
		}
	})

	getConfigMutex.Lock()

	return &getConfigMutex, &conf
}

func SetLevel(level logLevel) {
	lock, conf := getConfig()
	defer lock.Unlock()

	conf.level = level
}

func SetTimeFormat(format string) {
	lock, conf := getConfig()
	defer lock.Unlock()

	conf.timeFormat = format
}

func SetWriter(writer io.Writer) {
	lock, conf := getConfig()
	defer lock.Unlock()

	conf.writer = writer
}
