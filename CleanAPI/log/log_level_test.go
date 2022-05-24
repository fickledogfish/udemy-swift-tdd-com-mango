package log

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringShouldConvertValidLogLevelsToTheirCorrespondingStrings(t *testing.T) {
	// Arrange
	for _, c := range []struct {
		level          logLevel
		expectedString string
	}{
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarning, "WARN"},
		{LevelError, "ERR"},
	} {
		// Act
		var levelString string
		require.NotPanics(t, func() { levelString = c.level.String() })

		// Assert
		assert.Equal(t, c.expectedString, levelString)
	}
}

func TestStringShouldPanicOnUnexpectedLogLevels(t *testing.T) {
	// Arrange
	invalidLevel := logLevel(math.MaxUint)

	// Assert
	assert.PanicsWithValue(
		t,
		fmt.Sprintf("Unknown log level: %d", invalidLevel),
		func() {
			// Act
			_ = invalidLevel.String()
		})
}
