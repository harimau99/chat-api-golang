package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	db *sql.DB
}

// Generic method to handle common tasks such as logging
func (h *Handler) Process(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Start the stopwatch for logging the request
		start := time.Now()

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		// Call the handler and pass the original parameters
		handle(w, r, p)

		// Print the log to the console
		log.Printf(
			"%-6s%-20s%-20s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

// GET "/messages"
func (h *Handler) MessagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	messages, err := MessagesRetrieve(h.db)
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		log.Fatal(err)
	}
}

// POST "/messages"
func (h *Handler) MessagesPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Read body of request, but limit input to save server resources
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}

	// Unmarshal the json body into a Message struct
	var message Message
	if err := json.Unmarshal(body, &message); err != nil {
		// If error, return an HTTP status 422: Unprocessable entity
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	err = message.Create(h.db)
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

	w.WriteHeader(http.StatusCreated)
}
