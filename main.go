package main

import (
	"github.com/nathanfabio/api-usingGoland/internal/todo"
	"github.com/nathanfabio/api-usingGoland/internal/transport"
	"log"
)

func main() {

	svc := todo.NewService()
	server := transport.NewServer(svc)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
