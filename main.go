package main

import (
	"log"
	"net/http"
)

func main() {
	// See routes.go for a list of specified routes
	// See handlers.go for a list of controllers
	router := Router()

	log.Fatal(http.ListenAndServe(":8080", router))
}
