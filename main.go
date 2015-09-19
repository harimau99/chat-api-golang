package main

import (
	"log"
	"net/http"
)

func main() {
	// See routing.go
	router := Router()

	log.Fatal(http.ListenAndServe(":8080", router))
}
