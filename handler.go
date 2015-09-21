package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	db *sql.DB
}

// GET "/messages"
func (h *Handler) MessagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var messages Messages

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get the Messages from the database
	// Note: db.Query() opens and holds a connection until rows.Close()
	rows, err := h.db.Query("SELECT * FROM message ORDER BY created DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through query rows
	for rows.Next() {
		message := new(Message)

		// Scan gets the columns one row at a time
		err := rows.Scan(&message.Text, &message.Created)
		if err != nil {
			log.Fatal(err)
		}

		// Add the message to the Messages array
		messages = append(messages, *message)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		log.Fatal(err)
	}
}

// POST "/messages"
func (h *Handler) MessagesPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var message Message

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Read body of request, but limit input to save server resources
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}

	// Unmarshal the json body into a Message struct
	if err := json.Unmarshal(body, &message); err != nil {
		// If error, return an HTTP status 422: Unprocessable entity
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	// Create the message in the database
	stmt, err := h.db.Prepare("INSERT INTO message(created, text) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the statement with the data
	_, err = stmt.Exec(time.Now(), message.Text)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
