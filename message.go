package main

import (
	"time"
)

// Note the struct tags to ensure idiomatic JSON (lowercase keys)
type Message struct {
	Text    string    `json:"text"`
	Created time.Time `json:"created"`
}

type Messages []Message
