package signup

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ct "example.com/api/crypt/test"
	"example.com/api/models"
	"example.com/api/testutils"
	v "example.com/api/validations"
	vt "example.com/api/validations/test"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTPShouldRejectNonPOSTMethodsWithForbidden(t *testing.T) {
	for _, method := range []string{
		"GET",
		"PUT",
		"DELETE",
		"UPDATE",
	} {
		// Arrange
		sut := makeHandlerSut()
		sut.Request.Method = method

		// Act
		sut.ServeHTTP()

		// Assert
		assert.Equal(t, http.StatusForbidden, sut.ResponseRecorder.Code)
	}
}

func TestServeHTTPShouldSetContentTypeHeader(t *testing.T) {
	// Arrange
	sut := makeHandlerSut()

	// Act
	sut.ServeHTTP()

	// Assert
	assert.Contains(t, sut.ResponseRecorder.Header(), "Content-Type")
	assert.Contains(
		t,
		sut.ResponseRecorder.Header()["Content-Type"],
		"application/json; charset=utf-8",
	)
}

func TestServeHTTPShouldReturnInternalErrorOnBodyReaderFailure(t *testing.T) {
	// Arrange
	sut := makeHandlerSut()
	sut.Request.Body = testutils.NewReadCloserMock(
		func([]byte) (int, error) {
			return 0, errors.New("generic error")
		}, func() error {
			return nil
		},
	)

	// Act
	sut.ServeHTTP()

	// Assert
	assert.Equal(t, http.StatusInternalServerError, sut.ResponseRecorder.Code)
}

func TestServeHTTPShouldReturnBadRequestOnInvalidRequest(t *testing.T) {
	// Arrange
	sut := makeHandlerSut()
	sut.Request.Body = io.NopCloser(strings.NewReader(
		"{\"invalid_key\": \"invalid_value\"}",
	))

	// Act
	sut.ServeHTTP()

	// Assert
	assert.Equal(t, http.StatusBadRequest, sut.ResponseRecorder.Code)
}

// File SUT -------------------------------------------------------------------

type handlerSut struct {
	Request          *http.Request
	ResponseRecorder *httptest.ResponseRecorder

	ModelValidator *vt.ValidatorMock[models.SignUp]
	PasswordHasher *ct.PasswordHasherMock

	Handler handler
}

func makeHandlerSut() handlerSut {
	request := httptest.NewRequest("POST", "/", nil)
	responseRecorder := httptest.NewRecorder()

	modelValidator := vt.NewValidatorMock(func(models.SignUp) []v.Error {
		return []v.Error{}
	})

	passwordHasher := ct.NewPasswordHasherMock(func(string) ([]byte, error) {
		return []byte{}, nil
	})

	handler := handler{
		modelValidator: &modelValidator,
		passwordHasher: &passwordHasher,
	}

	return handlerSut{
		Request:          request,
		ResponseRecorder: responseRecorder,

		ModelValidator: &modelValidator,
		PasswordHasher: &passwordHasher,

		Handler: handler,
	}
}

func (sut *handlerSut) ServeHTTP() {
	sut.Handler.ServeHTTP(sut.ResponseRecorder, sut.Request)
}
