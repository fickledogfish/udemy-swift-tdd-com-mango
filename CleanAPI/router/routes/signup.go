package routes

import (
	"io/ioutil"
	"net/http"

	"example.com/api/router/models"
	res "example.com/api/router/responses"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method != http.MethodPost {
		res.Forbidden(w, "Forbidden")
		return
	}

	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.BadRequest(w, "Failed to read request")
		return
	}

	var reqAccountData models.SignUpModel
	if err = reqAccountData.UnmarshalBinary(bodyData); err != nil {
		res.InternalServerError(w, "Failed to parse request")
		return
	}

	if reqAccountData.Name == "" ||
		reqAccountData.Email == "" ||
		reqAccountData.Password == "" ||
		reqAccountData.PasswordConfirmation == "" {
		res.BadRequest(w, "Missing fields")
		return
	}

	if reqAccountData.Password != reqAccountData.PasswordConfirmation {
		res.BadRequest(w, "Password confirmation does not match")
		return
	}

	respModel := models.NewSignUpModelResponse(reqAccountData)

	res.Ok(w, respModel)
}
