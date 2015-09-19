package main

import (
	"time"
)

type Message struct {
	Text    string
	Created time.Time
}

type Messages []Message
