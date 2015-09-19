package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET "/messages"
func MessagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	messages := Messages{
		Message{Text: "Hi Oscar!"},
		Message{Text: "Hi Todd!"},
	}

	json.NewEncoder(w).Encode(messages)
}

// POST "/messages"
func MessagesPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "[post]\n")
}
