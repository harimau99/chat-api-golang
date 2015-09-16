package main

// Benchamrk suggest this router is fastest:
// https://github.com/julienschmidt/go-http-routing-benchmark
import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routing() {
	router := httprouter.New()

	router.GET("/messages", messagesGet)
	router.POST("/messages", messagesPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func messagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "[get]\n")
}

func messagesPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "[post]\n")
}
