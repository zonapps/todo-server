package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	log.Printf("Starting ToDo server\n")
	router := NewRouter()
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":60321", handler))
}
