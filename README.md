# chat-api-golang

## Setup

Follow this guide to setup Go on OS X: http://todsul.com/tech/setup-golang-on-mac-os-x/

Then install, build and run the project:

    git clone https://github.com/todsul/chat-api-golang.git
    cd chat-api-golang
    go get
    go install .
    chat-api-golang

For the mysql database

    User: root
    Password: [blank]
    Database: chat


    CREATE TABLE `message` (
      `text` text,
      `created` datetime DEFAULT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

## Testing

    go test ./tests

## API

    GET /messages (retrieve all messages)
    POST /messages (post new message)

## Coding

 To format code according to Go standards:

    go fmt

## Run

To list all messages:

    http://localhost:8080/messages

To post a message:

    curl -H "Content-Type: application/json" -d '{"text":"Hello"}' http://localhost:8080/messages

## TODO

* Move database operations to models (easier unit testing, simpler handlers, etc)
* Re-do handlers with centralized error handling and logging
* Write unit tests

## Reading

* https://golang.org/doc/code.html
* https://golang.org/doc/effective_go.html
* https://gobyexample.com/
