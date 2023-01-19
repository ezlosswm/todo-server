package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
	mux.HandleFunc("/todo", HandleTodo)
	mux.HandleFunc("/todo/{id}", GetTodoByID)

	//logs & runs server
	fmt.Printf("New API Server running on port %s\n", s.listenAddr)
	http.ListenAndServe(s.listenAddr, mux)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
