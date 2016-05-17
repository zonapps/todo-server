package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting ToDo server\n")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":60321", router))
}
