package main

import (
	"net/http"
	"web/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)

	_ = http.ListenAndServe(":8000", nil)
}
