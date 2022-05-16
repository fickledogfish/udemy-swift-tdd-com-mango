package responses

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnsureOkWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	message := "ok"
	response := NewDummyResponse(message)

	writer := httptest.NewRecorder()

	// Act
	Ok(writer, response)

	body, err := ioutil.ReadAll(writer.Body)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(
		t,
		fmt.Sprintf("{\"message\":\"%s\"}", message),
		string(body),
	)
}

func TestEnsureBadRequestWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	message := "bad request"
	writer := httptest.NewRecorder()

	// Act
	BadRequest(writer, message)

	body, err := ioutil.ReadAll(writer.Body)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusBadRequest, writer.Code)
	assert.Equal(
		t,
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		string(body),
	)
}

func TestEnsureInternalServerErrorWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	writer := httptest.NewRecorder()

	// Act
	InternalServerError(writer)

	body, err := ioutil.ReadAll(writer.Body)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, writer.Code)
	assert.Equal(
		t,
		"{\"error\":\"Internal server error\"}",
		string(body),
	)
}

func TestEnsureForbiddenWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	message := "forbidden"
	writer := httptest.NewRecorder()

	// Act
	Forbidden(writer, message)

	body, err := ioutil.ReadAll(writer.Body)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusForbidden, writer.Code)
	assert.Equal(
		t,
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		string(body),
	)
}

// Helper types ---------------------------------------------------------------

type DummyResponse struct {
	Message string `json:"message"`
}

func NewDummyResponse(message string) DummyResponse {
	return DummyResponse{
		Message: message,
	}
}

func (d DummyResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}
