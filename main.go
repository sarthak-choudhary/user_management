package main

import (
	"go_cassandra/api"
	"log"
	"net/http"
)

func main() {

	InitializeApp()

	http.Handle("/add_user", api.Adduser())
	log.Println("HTTP server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
