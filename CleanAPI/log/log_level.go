package log

import "fmt"

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
