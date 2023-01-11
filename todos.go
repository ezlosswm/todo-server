package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var listOfTodo TodoList

type TodoItem struct {
	ID          int       `json:"id"`
	Activity    string    `json:"activity"`
	CompletedAt time.Time `json:"completed_at"`
}

type TodoList struct {
	Items []TodoItem `json:"items"`
}

func NewTodo(activity string) *TodoList {
	return &TodoList{
		Items: []TodoItem{
			{
				ID:          rand.Intn(100),
				Activity:    activity,
				CompletedAt: time.Now().UTC()},
		},
	}
}

func (t *TodoList) GetItemByID(r *http.Request) (*TodoItem, error) {
	id, err := getID(r)
	if err != nil {
		return nil, err
	}

	for itemID := range listOfTodo.Items {
		if id == listOfTodo.Items[itemID].ID {
			return &listOfTodo.Items[itemID], nil
		}
	}

	return nil, fmt.Errorf("item with id %d not found", id)
}

func (t *TodoList) AddItem(item string) *TodoList {
	todoList := NewTodo(item)
	t.Items = append(t.Items, todoList.Items...)

	return todoList
}

func (t *TodoList) UpdateItem(item TodoList) *TodoList {
	return nil
}

func (t *TodoList) DeleteItemByID(item TodoList) *TodoList {
	return nil
}

func getID(r *http.Request) (int, error) {
	// retrieves the id from the post data
	idFromUser := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idFromUser)
	if err != nil {
		return id, fmt.Errorf("%d not a valid id", id)
	}

	return id, nil
}
