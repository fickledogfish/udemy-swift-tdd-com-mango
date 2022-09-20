package responses

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestResponsesSuite(t *testing.T) {
	suite.Run(t, new(responsesSuite))
}

type responsesSuite struct {
	suite.Suite

	ResponseRecorder *httptest.ResponseRecorder
	Response         dummyResponse
}

func (s *responsesSuite) SetupTest() {
	s.ResponseRecorder = httptest.NewRecorder()

	s.Response = dummyResponse{
		Message: "some_message",
		MarshalWith: func(d dummyResponse) ([]byte, error) {
			return json.Marshal(d)
		},
	}
}

// Test cases -----------------------------------------------------------------

func (s responsesSuite) TestWriteResponseShouldAnswerWithInternalErrorOnMarshalFailure() {
	// Arrange
	s.Response.MarshalWith = func(d dummyResponse) ([]byte, error) {
		return []byte{}, errors.New("error")
	}

	// Act
	writeResponse(s.ResponseRecorder, http.StatusOK, s.Response)

	// Assert
	s.Assert().Equal(
		"{\"error\":\"Internal server error\"}",
		s.readResponseBody(),
	)
}

func (s responsesSuite) TestEnsureOkWritesStatusHeaderAndResponse() {
	// Arrange
	message := "ok"

	s.Response.Message = message

	// Act
	Ok(s.ResponseRecorder, s.Response)

	// Assert
	s.Assert().Equal(http.StatusOK, s.ResponseRecorder.Code)
	s.Assert().Equal(
		fmt.Sprintf("{\"message\":\"%s\"}", message),
		s.readResponseBody(),
	)
}

func (s responsesSuite) TestEnsureBadRequestWritesStatusHeaderAndResponse() {
	// Arrange
	message := "bad request"

	s.Response.Message = message

	// Act
	BadRequest(s.ResponseRecorder, message)

	// Assert
	s.Assert().Equal(http.StatusBadRequest, s.ResponseRecorder.Code)
	s.Assert().Equal(
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		s.readResponseBody(),
	)
}

func (s responsesSuite) TestEnsureInternalServerErrorWritesStatusHeaderAndResponse() {
	// Act
	InternalServerError(s.ResponseRecorder)

	// Assert
	s.Assert().Equal(http.StatusInternalServerError, s.ResponseRecorder.Code)
	s.Assert().Equal(
		"{\"error\":\"Internal server error\"}",
		s.readResponseBody(),
	)
}

func (s responsesSuite) TestEnsureForbiddenWritesStatusHeaderAndResponse() {
	// Arrange
	message := "forbidden"

	s.Response.Message = message

	// Act
	Forbidden(s.ResponseRecorder, message)

	// Assert
	s.Assert().Equal(http.StatusForbidden, s.ResponseRecorder.Code)
	s.Assert().Equal(
		fmt.Sprintf("{\"error\":\"%s\"}", message),
		s.readResponseBody(),
	)
}

func (s responsesSuite) TestEnsureNotFoundWritesStatusHeaderAndResponse() {
	// Act
	NotFound(s.ResponseRecorder)

	// Assert
	s.Assert().Equal(http.StatusNotFound, s.ResponseRecorder.Code)
	s.Assert().Equal(
		"{\"error\":\"Not found\"}",
		s.readResponseBody(),
	)
}

// Helpers --------------------------------------------------------------------

func (s responsesSuite) readResponseBody() string {
	body, err := ioutil.ReadAll(s.ResponseRecorder.Body)
	s.Require().NoError(err)

	return string(body)
}

type dummyResponse struct {
	Message string `json:"message"`

	MarshalWith func(d dummyResponse) ([]byte, error) `json:"-"`
}

func (d dummyResponse) MarshalBinary() ([]byte, error) {
	return d.MarshalWith(d)
}
