package main

import (
	"board/globals"
	"board/handlers"
	"log"
	"net"
	"net/http"
)

func main() {
	defer globals.DB.Close()

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/post", handlers.Post)
	http.HandleFunc("/b/", handlers.Board)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started at localhost:8080")

	http.Serve(listener, nil)
}
