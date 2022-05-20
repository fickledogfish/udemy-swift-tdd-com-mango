package responses

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteResponseShouldAnswerWithInternalErrorOnMarshalFailure(t *testing.T) {
	// Arrange
	sut := makeResponsesSut()
	sut.Response.MarshalWith = func() ([]byte, error) {
		return []byte{}, errors.New("error")
	}

	// Act
	writeResponse(sut.ResponseRecorder, http.StatusOK, sut.Response)

	// Assert
	assert.Equal(
		t,
		"{\"error\":\"Internal server error\"}",
		sut.ReadResponseBody(t),
	)
}

func TestEnsureOkWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	message := "ok"

	sut := makeResponsesSut()
	sut.Response.Message = message

	// Act
	Ok(sut.ResponseRecorder, sut.Response)

	// Assert
	assert.Equal(t, http.StatusOK, sut.ResponseRecorder.Code)
	assert.Equal(
		t,
		fmt.Sprintf("{\"message\":\"%s\"}", message),
		sut.ReadResponseBody(t),
	)
}

func TestEnsureBadRequestWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	message := "bad request"

	sut := makeResponsesSut()
	sut.Response.Message = message

	// Act
	BadRequest(sut.ResponseRecorder, message)

	// Assert
	assert.Equal(t, http.StatusBadRequest, sut.ResponseRecorder.Code)
	assert.Equal(
		t,
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		sut.ReadResponseBody(t),
	)
}

func TestEnsureInternalServerErrorWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	sut := makeResponsesSut()

	// Act
	InternalServerError(sut.ResponseRecorder)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, sut.ResponseRecorder.Code)
	assert.Equal(
		t,
		"{\"error\":\"Internal server error\"}",
		sut.ReadResponseBody(t),
	)
}

func TestEnsureForbiddenWritesStatusHeaderAndResponse(t *testing.T) {
	// Arrange
	message := "forbidden"

	sut := makeResponsesSut()
	sut.Response.Message = message

	// Act
	Forbidden(sut.ResponseRecorder, message)

	// Assert
	assert.Equal(t, http.StatusForbidden, sut.ResponseRecorder.Code)
	assert.Equal(
		t,
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		sut.ReadResponseBody(t),
	)
}

// Helper types ---------------------------------------------------------------

type DummyResponse struct {
	Message string `json:"message"`

	MarshalWith func() ([]byte, error) `json:"-"`
}

func (d DummyResponse) MarshalBinary() ([]byte, error) {
	if d.MarshalWith != nil {
		return d.MarshalWith()
	}

	return json.Marshal(d)
}

// File SUT -------------------------------------------------------------------

type responsesSut struct {
	ResponseRecorder *httptest.ResponseRecorder
	Response         DummyResponse
}

func makeResponsesSut() responsesSut {
	return responsesSut{
		ResponseRecorder: httptest.NewRecorder(),
		Response: DummyResponse{
			Message: "some_message",
		},
	}
}

func (r responsesSut) ReadResponseBody(t *testing.T) string {
	body, err := ioutil.ReadAll(r.ResponseRecorder.Body)
	require.NoError(t, err)

	return string(body)
}
