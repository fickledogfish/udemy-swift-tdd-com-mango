package utils

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBodyShouldLimitReadSize(t *testing.T) {
	// Arrange
	limit := 7
	repeating := "a"

	body := strings.Repeat(repeating, limit*2)
	req := httptest.NewRequest("", "/", strings.NewReader(body))

	// Act
	data, err := ReadBody(req, int64(limit))

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, len(data), limit)
	assert.Equal(t, string(data), strings.Repeat(repeating, limit))
}
