package router

import (
	"net/http"

	"example.com/api/router/routes"
)

const HttpAddr = ":5050"

func NewRouter() http.Handler {
	// Subrouters

	apiMux := http.NewServeMux()
	apiMux.Handle("/signup", http.HandlerFunc(routes.Signup))
	apiMux.Handle("/login", http.HandlerFunc(routes.Login))

	// Root router

	rootMux := http.NewServeMux()
	rootMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	return rootMux
}
