package signup

import (
	"io"
	"io/ioutil"
	"net/http"

	"example.com/api/crypt"
	"example.com/api/log"
	"example.com/api/models"
	r "example.com/api/router/responses"
	u "example.com/api/router/routes/utils"
	v "example.com/api/validations"
	vm "example.com/api/validations/models"
)

type handler struct {
	modelValidator v.Validation[models.SignUp]
	passwordHasher crypt.IPasswordHasher
}

// Creates the default handler with its dependencies.
func NewHandler() http.Handler {
	modelValidator := vm.NewSignUpModelValidator(v.NewEmailValidator())
	passwordHasher := crypt.NewPasswordHasher()

	return NewHandlerWithOptions(
		modelValidator,
		passwordHasher,
	)
}

// Like NewHandler(), but this allows for the dependencies to be injected into
// the handler.
func NewHandlerWithOptions(
	modelValidator v.Validation[models.SignUp],
	passwordHasher crypt.IPasswordHasher,
) http.Handler {
	return handler{
		modelValidator: modelValidator,
		passwordHasher: passwordHasher,
	}
}

// Implementing http.Handler --------------------------------------------------

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	u.SetContentTypeHeader(w, u.ContentTypeApplicationJSON)

	if req.Method != http.MethodPost {
		r.Forbidden(w, "Forbidden")
		log.Info("%v", req.Method)
		return
	}

	bodyData, err := ioutil.ReadAll(io.LimitReader(req.Body, 32))
	if err != nil {
		r.BadRequest(w, "Failed to read request")
		log.Info(err.Error())
		return
	}

	var reqAccountData models.SignUp
	if err = reqAccountData.UnmarshalBinary(bodyData); err != nil {
		r.InternalServerError(w)
		return
	}

	if reqAccountData.Name == "" ||
		reqAccountData.Email == "" ||
		reqAccountData.Password == "" ||
		reqAccountData.PasswordConfirmation == "" {
		r.BadRequest(w, "Missing fields")
		return
	}

	if reqAccountData.Password != reqAccountData.PasswordConfirmation {
		r.BadRequest(w, "Password confirmation does not match")
		return
	}

	newUser, err := models.NewUser(h.passwordHasher, reqAccountData)
	if err != nil {
		r.InternalServerError(w)
		return
	}

	respModel := models.NewSignUpResponse(newUser)

	log.Debug("created: %+v", respModel)
	r.Ok(w, respModel)
}
