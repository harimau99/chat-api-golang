package main

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

func (m *Message) Create(db *sql.DB) (error) {
	// Create the message in the database
	stmt, err := db.Prepare("INSERT INTO message(created, text) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the statement with the data
	_, err = stmt.Exec(time.Now(), m.Text)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func MessagesRetrieve(db *sql.DB) (Messages, error) {
	var messages Messages

	// Get the Messages from the database
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

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return messages, nil
}
