package routes

import (
	"io/ioutil"
	"net/http"

	"example.com/api/log"
	"example.com/api/models"
	r "example.com/api/router/responses"
	v "example.com/api/validations"
	vm "example.com/api/validations/models"
)

type signUpHandler struct {
	modelValidator v.Validation[models.SignUp]
}

func NewSignUpHandler() signUpHandler {
	modelValidator := vm.NewSignUpModelValidator(v.NewEmailValidator())

	return NewSignUpHandlerWithOptions(
		modelValidator,
	)
}

func NewSignUpHandlerWithOptions(
	modelValidator v.Validation[models.SignUp],
) signUpHandler {
	return signUpHandler{
		modelValidator: modelValidator,
	}
}

// Implementing http.Handler --------------------------------------------------

func (h signUpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != http.MethodPost {
		r.Forbidden(w, "Forbidden")
		log.Info("%v", req.Method)
		return
	}

	bodyData, err := ioutil.ReadAll(req.Body)
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

	newUser, err := models.NewUser(reqAccountData)
	if err != nil {
		r.InternalServerError(w)
		return
	}

	respModel := models.NewSignUpResponse(newUser)

	log.Debug("created: %+v", respModel)
	r.Ok(w, respModel)
}
