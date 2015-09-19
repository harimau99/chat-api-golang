package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET "/messages"
func MessagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	messages := RetrieveMessages()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

// POST "/messages"
func MessagesPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	message := CreateMessage("Testing")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
