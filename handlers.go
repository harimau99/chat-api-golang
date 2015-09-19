package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// GET "/messages"
func MessagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	messages := Messages{
		Message{Text: "Hi Oscar!", Created: time.Now()},
		Message{Text: "Hi Todd!", Created: time.Now()},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
        panic(err)
    }
}

// POST "/messages"
func MessagesPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode("OK"); err != nil {
        panic(err)
    }
}
