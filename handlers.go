package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetTodo(w, r)
	case "POST":
		PostTodo(w, r)
	}
}

// GET methods
func GetTodo(w http.ResponseWriter, r *http.Request) {

	WriteJson(w, http.StatusOK, listOfTodo)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	// id := new(TodoItem)

	item, err := listOfTodo.GetItemByID(r)
	if err != nil {
		log.Println(err)
	}

	WriteJson(w, http.StatusOK, item)
}

// POST methods
func PostTodo(w http.ResponseWriter, r *http.Request) {
	item := new(TodoItem)

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatal(err)
	}

	listOfTodo.AddItem(item.Activity)

	WriteJson(w, http.StatusOK, listOfTodo)
}
