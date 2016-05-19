package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

var DatabaseConnection *sql.DB

func main() {
	dbFile := "/todo.db"

	if len(os.Args) > 1 {
		dbFile = os.Args[1]
	}
	log.Printf("Starting server on port 60321, sqlite DB file: %s\n", dbFile)

	err := ConnectToSqlite(dbFile)
	if err != nil {
		log.Fatalf("Error connecting to %s: %s\n", dbFile, err.Error())
	}
	router := NewRouter()

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":60321", handler))
}

func ConnectToSqlite(dbFile string) error {
	if DatabaseConnection != nil {
		DatabaseConnection.Close()
		DatabaseConnection = nil
	}

	var err error

	DatabaseConnection, err = sql.Open("sqlite3", dbFile)
	return err
}
