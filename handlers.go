package main

import (
	"net/http"

	"encoding/json"
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
	return []todoEntry{
		todoEntry{
			Title:   "Work on Todo List Server",
			DueDate: "2016-05-17T23:59:59-08:00",
			Description: "Give the Todo App functionality " +
				"to query from the server",
		},
		todoEntry{
			Title:       "Finish Work on Todo List UI",
			DueDate:     "2016-05-15T23:59:59-08:00",
			Description: `This is overdue man....`,
		},
		todoEntry{
			Title:       "Get new house",
			DueDate:     "2016-05-31T12:00:00-08:00",
			Description: `So excited man`,
		},
	}
}
