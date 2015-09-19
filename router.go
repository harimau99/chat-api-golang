package main

// Benchmarks suggest this router is fastest:
// https://github.com/julienschmidt/go-http-routing-benchmark
import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Function)
	}

	return router
}
