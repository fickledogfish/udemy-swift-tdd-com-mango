package router

import (
	"fmt"
	"net/http"

	"example.com/api/router/models"
)

func badRequest(w http.ResponseWriter, message string) {
	writeError(w, http.StatusForbidden, message)
}

func forbidden(w http.ResponseWriter, message string) {
	writeError(w, http.StatusForbidden, message)
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	fmt.Fprint(w, models.ReturnError{
		Message: message,
		Code:    code,
	})
}
