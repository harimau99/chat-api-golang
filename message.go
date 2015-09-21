package main

import (
	"time"
)

// Note the struct tags to ensure idiomatic JSON (lowercase keys)
type Message struct {
	Created time.Time `json:"created"`
	Text    string    `json:"text"`
}

type Messages []Message
