package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// handles requests made to /todo
func (s *APIServer) HandleTodo(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.HandleGetTodo(w, r)
	case "POST":
		return s.HandlePostTodo(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handles requests mde to todo/{id}
func (s *APIServer) HandleTodoByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.HandleGetTodoByID(w, r)
	case "DELETE":
		return s.HandleDeleteTodo(w, r)
	case "PUT":
		return s.HandleUpdateTodo(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) HandleGetTodo(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.store.GetTodos()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, todos)
}

func (s *APIServer) HandleGetTodoByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	item, err := s.store.GetTodoByID(id)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, map[string]int{"id does not exist": id})

		// return err
	}

	return WriteJson(w, http.StatusOK, item)
}

func (s *APIServer) HandlePostTodo(w http.ResponseWriter, r *http.Request) error {
	req := new(TodoReq)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	todo := NewTodo(req.Item, req.CreatedAt)

	err := s.store.CreateTodo(todo)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, todo)
}

func (s *APIServer) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err = s.store.DeleteTodo(id); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, map[string]int{"deleted item:": id})
}

func (s *APIServer) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
	newRequest := new(TodoList)

	if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
		return err
	}
	defer r.Body.Close()

	id, err := getID(r)
	if err != nil {
		return err
	}

	if err = s.store.UpdateTodo(id); err != nil {
		return WriteJson(w, http.StatusBadRequest, newRequest)
	}

	return WriteJson(w, http.StatusOK, map[int]string{id: "marked as completed"})
}

// reads the id input field and converts to int
// 400 represents invalid id
func getID(r *http.Request) (int, error) {
	idFromUser := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idFromUser)
	if err != nil {
		return 404, fmt.Errorf("invalid id: %v", id)
	}

	return id, nil
}
