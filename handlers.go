package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetTodo(w, r)
	case "POST":
		PostTodo(w, r)
	}
}

// GET methods
func GetTodo(w http.ResponseWriter, r *http.Request) {

	WriteJson(w, http.StatusOK, ListOfTodo)
}

// POST methods
func PostTodo(w http.ResponseWriter, r *http.Request) {
	item := new(TodoItem)

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatal(err)
	}

	ListOfTodo.AddItem(item.Activity)

	WriteJson(w, http.StatusOK, ListOfTodo)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	// saves the id as a local variable
	id, err := GetID(r)
	if err != nil {
		log.Println(err)
	}

	if r.Method == "GET" {
		item, err := ListOfTodo.GetItemByID(id, ListOfTodo)
		if err != nil {
			log.Println(err)

		}
		WriteJson(w, http.StatusOK, item)
	}

	if r.Method == "DELETE" {
		if err := ListOfTodo.DeleteItem(id, ListOfTodo); err != nil {
			log.Println(err)
		}
	}
}
