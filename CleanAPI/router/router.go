package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/api/router/models"
)

const HttpAddr = ":5050"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api", http.HandlerFunc(signinHandler))

	return mux
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method != http.MethodPost {
		forbidden(w, "Forbidden")
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var reqAccountData models.SignUpModel
	if err := decoder.Decode(&reqAccountData); err != nil {
		badRequest(w, "Failed to parse request")
		return
	}

	if reqAccountData.Name == "" ||
		reqAccountData.Email == "" ||
		reqAccountData.Password == "" ||
		reqAccountData.PasswordConfirmation == "" {
		badRequest(w, "Missing fields")
		return
	}

	if reqAccountData.Password != reqAccountData.PasswordConfirmation {
		badRequest(w, "Password confirmation does not match")
		return
	}

	fmt.Fprint(w, models.NewSignUpModelResponse(reqAccountData))
}
