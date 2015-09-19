package main

// Then underscore in front of an import does something I can't remember
import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Note the struct tags to ensure idiomatic JSON (lowercase keys)
type Message struct {
	Created time.Time `json:"created"`
	Text    string    `json:"text"`
}

type Messages []Message

// Maintaining CRUD naming convention for model functions
// TODO: abstract database code from here (maybe)

func CreateMessage(message Message) {
	// Get a mysql connection
	// To scan mysql.DateTime into time.Time, we need ?parseTime=true
	db, err := sql.Open("mysql", "root:@/chat")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO message(created, text) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the statement with the data
	res, err := stmt.Exec(message.Created, message.Text)
	if err != nil {
		log.Fatal(err)
	}

	// Log the details of this create
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(
		"%-6s%-20s%-20d",
		"SQL",
		"-create",
		lastId,
	)
}

func RetrieveMessages() Messages {
	var messages Messages

	// Get a mysql connection
	// To scan mysql.DateTime into time.Time, we need ?parseTime=true
	db, err := sql.Open("mysql", "root:@/chat?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get the Messages from the database
	// Yay, raw sql! We're together again at last!
	// Note: db.Query() opens and holds a connection until rows.Close()
	rows, err := db.Query("SELECT * FROM message ORDER BY created DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through query rows
	for rows.Next() {
		message := new(Message)

		// Scan gets the columns one row at a time
		err := rows.Scan(&message.Text, &message.Created)
		if err != nil {
			log.Fatal(err)
		}

		// Add the message to the Messages array
		messages = append(messages, *message)
	}

	// Log details of the retrieve
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(
		"%-6s%-20s%-20d",
		"SQL",
		"-retrieve",
		len(messages),
	)

	return messages
}
