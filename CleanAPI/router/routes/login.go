package routes

import "net/http"

type logInHandler struct {
}

func NewLogInHandler() logInHandler {
	return logInHandler{}
}

// Implementing http.Handler --------------------------------------------------

func (h logInHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}
