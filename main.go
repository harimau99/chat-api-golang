package main

import ()

func main() {
	// We must setup the database only once for efficiency with connections
	// See this page for information: http://go-database-sql.org/accessing.html
	db := Database()

	// See routes.go for a list of specified routes
	// See handlers.go for a list of handlers/controllers
	Router(db)
}
