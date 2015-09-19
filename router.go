package main

// Benchmarks suggest this router is fastest:
// https://github.com/julienschmidt/go-http-routing-benchmark
import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {
		// Wrap handler in logger function to log every request
        handler := Logger(route.Function, route.Name)

		// Pass the wrapped handler to the router
		router.Handle(route.Method, route.Pattern, handler)
	}

	return router
}
