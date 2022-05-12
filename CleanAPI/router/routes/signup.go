package routes

import (
	"io/ioutil"
	"net/http"

	"example.com/api/log"
	"example.com/api/router/models"
	r "example.com/api/router/responses"
)

func Signup(w http.ResponseWriter, req *http.Request) {
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

	var reqAccountData models.SignUpModel
	if err = reqAccountData.UnmarshalBinary(bodyData); err != nil {
		r.InternalServerError(w, "Failed to parse request")
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

	respModel := models.NewSignUpModelResponse(reqAccountData)

	log.Debug("created: %+v", respModel)
	r.Ok(w, respModel)
}
