package main

import (
	"net/http"

	"example.com/api/router"
)

func main() {
	http.ListenAndServe(router.HttpAddr, router.NewRouter())
}
