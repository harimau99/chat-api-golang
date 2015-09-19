package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Logger(handler httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Start the stopwatch for logging the request
		start := time.Now()

		// Call the handler and pass the original parameters
		handler(w, r, p)

		// Print the log to the console
		log.Printf(
			"%-6s%-20s%-20s%-20s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	}
}
