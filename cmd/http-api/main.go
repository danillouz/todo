package main

import (
	"log"

	"github.com/danillouz/todo/pkg/http/rest"
	"github.com/danillouz/todo/pkg/storage/memory"
)

func main() {
	service := memory.NewStorage()
	server := rest.NewServer(service)
	err := server.Run(":8888")

	if err != nil {
		log.Fatal(err)
	}
}
