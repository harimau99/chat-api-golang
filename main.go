package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/todsul/chat-api-golang/handlers"
	"github.com/todsul/chat-api-golang/models"
)

func main() {
	// We must setup the database only once for efficiency with connections
	// See this page for information: http://go-database-sql.org/accessing.html
	db := models.New()
	router := httprouter.New()
	handler := &handlers.Handler{db}

	// We wrap the handler to execute boilerplate handler code
	router.Handle("GET", "/messages", handler.Process(handler.MessagesGet))
	router.Handle("POST", "/messages", handler.Process(handler.MessagesPost))

	log.Fatal(http.ListenAndServe(":8080", router))
}
