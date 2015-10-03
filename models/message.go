package models

import (
	"time"
)

// Note the struct tags to ensure idiomatic JSON (lowercase keys)
type Message struct {
	Created time.Time `json:"created"`
	Text    string    `json:"text"`
}

type Messages []Message

func (db *DB) MessageCreate(text string) (error) {
	// Create the message in the database
	stmt, err := db.Prepare("INSERT INTO message(created, text) VALUES(?, ?)")
	if err != nil {
		return err
	}

	// Execute the statement with the data
	_, err = stmt.Exec(time.Now(), text)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) MessagesRetrieve() (Messages, error) {
	var messages Messages

	// Get the Messages from the database
	// Note: db.Query() opens and holds a connection until rows.Close()
	rows, err := db.Query("SELECT * FROM message ORDER BY created DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through query rows
	for rows.Next() {
		message := new(Message)

		// Scan gets the columns one row at a time
		err := rows.Scan(&message.Text, &message.Created)
		if err != nil {
			return nil, err
		}

		// Add the message to the Messages array
		messages = append(messages, *message)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return messages, nil
}
