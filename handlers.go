package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// handles general GET and POST requests
func HandleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetTodo(w, r)
	case "POST":
		PostTodo(w, r)
	}
}

// handles all methods that reacts to a specified todo item
func HandleTodoByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetTodoByID(w, r)
	case "DELETE":
		DeleteTodo(w, r)
	case "PUT":
		UpdateTodo(w, r)
	case "PATCH":
		CompleteTodo(w, r)
	}
}

// handler func for retreiving all items
func GetTodo(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, http.StatusOK, ListOfTodo)
}

// handler func returns a todo item specified by id
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(r)
	CheckErr(err)

	item, err := ListOfTodo.GetItemByID(id)
	CheckErr(err)

	WriteJson(w, http.StatusOK, item)
}

// POST methods
func PostTodo(w http.ResponseWriter, r *http.Request) {
	item := new(TodoItem)

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Print(err)
	}

	ListOfTodo.AddItem(item.Activity)

	WriteJson(w, http.StatusOK, ListOfTodo)
}

// allows edits to be made to a specifed todo item
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	updatedTodo := new(TodoItem)

	json.NewDecoder(r.Body).Decode(&updatedTodo)
	defer r.Body.Close()

	id, err := GetID(r)
	CheckErr(err)

	newTodoList, err := ListOfTodo.UpdateItem(id, *updatedTodo)
	CheckErr(err)

	WriteJson(w, http.StatusOK, newTodoList)
}

// makes todo items as completed
func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&ListOfTodo)
	defer r.Body.Close()

	id, err := GetID(r)
	CheckErr(err)

	completed, err := ListOfTodo.CompleteItem(id, ListOfTodo.Items)
	CheckErr(err)

	WriteJson(w, http.StatusOK, completed)
}

// deletes a todo item based on specified id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(r)
	CheckErr(err)

	newTodos, err := ListOfTodo.DeleteItem(id)
	CheckErr(err)

	WriteJson(w, http.StatusOK, newTodos)
}
