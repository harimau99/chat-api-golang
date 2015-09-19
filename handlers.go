package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

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
	var message Message

	// Read body of request, but limit input to save server resources
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Unmarshal the json body into a Message struct
	// How beautifully convenient?!
	if err := json.Unmarshal(body, &message); err != nil {
		// If error, return an HTTP status 422: Unprocessable entity
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422)

        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

	// Create the message in the database
	message.Created = time.Now()
	CreateMessage(message)

	// Return an HTTP status CREATED
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
