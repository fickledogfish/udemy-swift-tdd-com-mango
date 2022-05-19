package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetContentTypeHeaderShouldSetTheContentTypeHeader(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()

	// Act
	SetContentTypeHeader(w, ContentTypeApplicationJSON)

	// Assert
	assert.Contains(t, w.Header(), "Content-Type")
}
