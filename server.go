package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var DatabaseConnection *sql.DB

func main() {
	log.Printf("Testing123\n")

	err := ConnectToSqlite("/todo.db")
	if err != nil {
		log.Fatalf("Error connecting to /todo.db: %s\n", err.Error())
	}
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":60321", router))
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
