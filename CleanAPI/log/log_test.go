package log

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestLogTestSuite(t *testing.T) {
	suite.Run(t, new(logTestSuite))
}

type logTestSuite struct {
	suite.Suite

	level      logLevel
	timeFormat string

	strBuilder *strings.Builder
}

func (l *logTestSuite) SetupTest() {
	l.strBuilder = &strings.Builder{}
	l.level = LevelDebug
	l.timeFormat = "time_format"

	SetLevel(l.level)
	SetTimeFormat(l.timeFormat)
	SetWriter(l.strBuilder)
}

func (l *logTestSuite) TearDownTest() {
}

func (l *logTestSuite) TestWriteLogShouldWriteTheMessageStringToAChannel() {
	// Arrange
	level := l.level
	message := "some_message"

	ch := make(chan string, 1)

	// Act
	writeLog(ch, level, message)
	l.Require().Eventually(
		func() bool { return len(ch) > 0 },
		1*time.Second,
		100*time.Millisecond,
	)

	// Assert
	l.Equal(
		fmt.Sprintf("[%s] %s => %s", level, l.timeFormat, message),
		<-ch,
	)
}
