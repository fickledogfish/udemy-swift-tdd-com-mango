package log

import (
	"sync"
	"time"
)

type logConfig struct {
	level      logLevel
	timeFormat string
}

var conf logConfig

var initLogConfig sync.Once

func getConfig() *logConfig {
	initLogConfig.Do(func() {
		conf = logConfig{
			level:      LevelDebug,
			timeFormat: time.RFC3339,
		}
	})

	return &conf
}
