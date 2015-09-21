package main

// Benchmarks suggest this router is fastest:
// https://github.com/julienschmidt/go-http-routing-benchmark
import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router(db *sql.DB) {
	router := httprouter.New()

	// We use a handler struct to pass hold a single db instance
	handler := Handler{db: db}

	// We wrap the handler with a logger to output to console
	router.Handle("GET", "/messages", HandleLogger(handler.MessagesGet))
	router.Handle("POST", "/messages", HandleLogger(handler.MessagesPost))

	log.Fatal(http.ListenAndServe(":8080", router))
}
