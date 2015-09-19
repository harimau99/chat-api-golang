package main

var routes = Routes{
	Route{
		"MessagesGet",
		"GET",
		"/messages",
		MessagesGet,
	},
	Route{
		"MessagesPost",
		"POST",
		"/messages",
		MessagesPost,
	},
}
