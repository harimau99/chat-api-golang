package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

func main() {
    router := httprouter.New()
    router.GET("/messages", messagesGet)
    router.POST("/messages", messagesPost)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func messagesGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "[list messages]\n")
}

func messagesPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "[posted message]\n")
}
