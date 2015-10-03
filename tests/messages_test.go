package tests

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/julienschmidt/httprouter"

    "github.com/todsul/chat-api-golang/handlers"
    "github.com/todsul/chat-api-golang/mocks"
)

func TestMessagesGet(t *testing.T) {
    rec := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/messages", nil)

    // DB is mocked
    router := httprouter.New()
    handler := &handlers.Handler{DB: &mocks.MockDB{}}

    router.Handle("GET", "/messages", handler.Process(handler.MessagesGet))
    router.ServeHTTP(rec, req)

    expected := "[{\"created\":\"2001-01-01T01:00:00Z\",\"text\":\"Hello\"},{\"created\":\"2001-01-01T01:00:00Z\",\"text\":\"Howdy\"}]\n"
    if expected != rec.Body.String() {
        t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
    }
}

func TestMessagesPost(t *testing.T) {
    rec := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/messages", bytes.NewBufferString("Hello"))

    // DB is mocked
    router := httprouter.New()
    handler := &handlers.Handler{DB: &mocks.MockDB{}}

    router.Handle("POST", "/messages", handler.Process(handler.MessagesPost))
    router.ServeHTTP(rec, req)

    expected := "\n"
    if expected != rec.Body.String() {
        t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
    }
}
