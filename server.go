package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Testing123\n")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":60321", router))
}
