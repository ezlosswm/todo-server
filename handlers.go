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

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// var updatedTodo TodoItem
	updatedTodo := new(TodoItem)
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	defer r.Body.Close()

	id, err := GetID(r)
	if err != nil {
		log.Print(err)
	}

	l := ListOfTodo.UpdateItem(id, *updatedTodo)
	WriteJson(w, http.StatusOK, l)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&ListOfTodo)
	defer r.Body.Close()

	id, err := GetID(r)
	if err != nil {
		log.Print(err)
	}

	WriteJson(w, http.StatusOK, ListOfTodo.CompleteItem(id, ListOfTodo.Items))
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	// saves the id as a local variable
	id, err := GetID(r)
	if err != nil {
		log.Println(err)
	}

	if r.Method == "GET" {
		item, err := ListOfTodo.GetItemByID(id)
		if err != nil {
			log.Println(err)

		}
		WriteJson(w, http.StatusOK, item)
	}

	if r.Method == "DELETE" {
		if err := ListOfTodo.DeleteItem(id); err != nil {
			log.Println(err)
		}
	}

	if r.Method == "PUT" {
		UpdateTodo(w, r)
	}

	if r.Method == "PATCH" {
		CompleteTodo(w, r)
	}
}
