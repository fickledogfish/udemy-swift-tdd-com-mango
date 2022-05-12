package main

import (
	"fmt"
	"net/http"

	"example.com/api/router"
)

func main() {
	err := http.ListenAndServe(router.HttpAddr, router.NewRouter())
	if err != nil {
		fmt.Println(err)
	}
}
