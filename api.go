package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenAddr string
	store      Storager
}

func NewAPIServer(listenAddr int, store Storager) *APIServer {
	return &APIServer{
		ListenAddr: fmt.Sprintf(":%d", listenAddr),
		store:      store,
	}
}

func (s *APIServer) Run() {
	// starts new router
	mux := mux.NewRouter()

	// handler functions
	mux.HandleFunc("/todo", makeHandlerFunc(s.HandleTodo))
	mux.HandleFunc("/todo/{id}", makeHandlerFunc(s.HandleTodoByID))

	//logs & runs server
	fmt.Printf("New API Server running on port %s\n", s.ListenAddr)
	http.ListenAndServe(s.ListenAddr, mux)
}

// adds head & encodes json
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// holds handle function signature
type apiHandler func(http.ResponseWriter, *http.Request) error

// converts apiHandler to http.HandlerFunc
func makeHandlerFunc(h apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, err)
		}
	}
}
