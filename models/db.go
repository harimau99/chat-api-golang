package models

// Then underscore in front of an import does something I can't remember
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
    MessageCreate(string) (error)
	MessagesRetrieve() (Messages, error)
}

type DB struct {
    *sql.DB
}

func New() (*DB) {
	db, err := sql.Open("mysql", "root:@/chat?parseTime=true")
	if err != nil {
	    log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
	    log.Fatal(err)
	}
	return &DB{db}
}
