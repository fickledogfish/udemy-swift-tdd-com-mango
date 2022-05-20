package signup

import (
	"net/http"
	"net/http/httptest"
	"testing"

	ct "example.com/api/crypt/test"
	"example.com/api/models"
	v "example.com/api/validations"
	vt "example.com/api/validations/test"
	"github.com/stretchr/testify/assert"
)

func TestServeHTTPShouldSetContentTypeHeader(t *testing.T) {
	// Arrange
	sut := makeHandlerSut()

	// Act
	sut.Handler.ServeHTTP(sut.ResponseRecorder, sut.Request)

	// Assert
	assert.Contains(t, sut.ResponseRecorder.Header(), "Content-Type")
	assert.Contains(
		t,
		sut.ResponseRecorder.Header()["Content-Type"],
		"application/json; charset=utf-8",
	)
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
	request := httptest.NewRequest("GET", "/", nil)
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
