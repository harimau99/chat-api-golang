package handlers

import (
    "encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

    "github.com/julienschmidt/httprouter"

    "github.com/todsul/chat-api-golang/models"
)

// GET "/messages"
func (h *Handler) MessagesGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	messages, err := h.DB.MessagesRetrieve()
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
	var message models.Message
	if err := json.Unmarshal(body, &message); err != nil {
		// If error, return an HTTP StatusNotAcceptable
		w.WriteHeader(http.StatusNotAcceptable)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
		return
	}

    // Create the message on the database
    if err = h.DB.MessageCreate(message.Text); err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
}
