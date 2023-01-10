package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var listOfTodo TodoList

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr int) *APIServer {
	return &APIServer{
		listenAddr: fmt.Sprintf(":%d", listenAddr),
	}
}

func (s *APIServer) Run() {
	// starts new router
	mux := mux.NewRouter()

	// handler functions
	mux.HandleFunc("/todo", handleTodo)

	//logs & runs server
	fmt.Printf("New API Server running on port %s\n", s.listenAddr)
	http.ListenAndServe(s.listenAddr, mux)
}

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

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
