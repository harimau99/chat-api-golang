package mocks

import (
    "time"

    "github.com/todsul/chat-api-golang/models"
)

type MockDB struct{}

func (m *MockDB) MessagesRetrieve() (models.Messages, error) {
    var messages models.Messages

    created, _ := time.Parse(time.RFC3339, "2001-01-01T01:00:00+00:00")

    messages = append(messages, models.Message{created, "Hello"})
    messages = append(messages, models.Message{created, "Howdy"})

    return messages, nil
}

func (m *MockDB) MessageCreate(text string) error {
    var message models.Message

    message.Text = text

    return nil
}
