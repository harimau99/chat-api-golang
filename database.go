package main

// Then underscore in front of an import does something I can't remember
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root:@/chat?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
