package main

import (
	"log"
	"net/http"
	"time"

	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

type todoEntry struct {
	Title       string `json:"title"`
	DueDate     string `json:"date_due"`
	Description string `json:"description"`
}

func TodoEntries(w http.ResponseWriter, r *http.Request) {
	todoEntries := getTodoItems()

	js, err := json.Marshal(todoEntries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getTodoItems() []todoEntry {
	rows, err := DatabaseConnection.Query("SELECT title,description,date_due from todo_item")
	if err != nil {
		log.Printf("Error querying rows: %s\n", err.Error())
		return nil
	}

	var entries []todoEntry
	var entry todoEntry
	var unixTime int64 = 0
	for rows.Next() {
		err = rows.Scan(&entry.Title, &entry.Description, &unixTime)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return entries
		}
		tm := time.Unix(unixTime, 0)
		date := tm.Format("2006-01-_2T15:04:05-08:00")
		entry.DueDate = date
		entries = append(entries, entry)

	}

	return entries
}
