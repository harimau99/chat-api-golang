package main

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name     string
	Method   string
	Pattern  string
	Function httprouter.Handle
}

type Routes []Route
