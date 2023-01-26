package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// creates a new todo item within a todo list
func NewTodo(activity string) *TodoList {
	return &TodoList{
		Items: []TodoItem{
			{
				ID:          rand.Intn(100),
				Activity:    activity,
				CompletedAt: time.Time{},
				Completed:   false},
		},
	}
}

// retrieves todo item by the specified id
func (t *TodoList) GetItemByID(id int) (*TodoItem, error) {
	for itemID := range t.Items {
		if id == t.Items[itemID].ID {
			return &t.Items[itemID], nil
		}
	}

	return nil, fmt.Errorf("item with id %d not found", id)
}

// appends new item to the current list of todos
func (t *TodoList) AddItem(item string) *TodoList {
	todoList := NewTodo(item)
	t.Items = append(t.Items, todoList.Items...)

	return todoList
}

// allows edits to be made to the todo item activity
func (t *TodoList) UpdateItem(id int, newItem TodoItem) (*TodoList, error) {
	for i, todo := range t.Items {
		if todo.ID == id {

			newItem.ID = id
			todo.ID = newItem.ID
			todo.Activity = newItem.Activity
			// todo.CompletedAt = newItem.CompletedAt
			// todo.Completed = newItem.Completed

			t.Items[i] = newItem

			return t, nil
		}
	}

	return nil, fmt.Errorf("unable to update %v", t.Items)
}

// logs the time marked as completed
func (t *TodoList) CompleteItem(id int, completedItem []TodoItem) (*TodoList, error) {
	for i, todo := range t.Items {
		if todo.ID == id {
			completedItem[i].CompletedAt = time.Now()
			completedItem[i].Completed = true

			todo.CompletedAt = completedItem[i].CompletedAt
			todo.Completed = completedItem[i].Completed

			return t, nil
		}
	}

	return nil, fmt.Errorf("unable to %v mark as completed", t.Items)
}

func (t *TodoList) DeleteItem(id int) (*TodoList, error) {
	for i, todo := range t.Items {
		if todo.ID == id {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
		}
	}

	return t, fmt.Errorf("id %s not found", strconv.Itoa(id))
}

// reads the id input field and converts to an integer
func GetID(r *http.Request) (int, error) {
	idFromUser := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idFromUser)
	// CheckErr(err)
	if err != nil {
		return id, fmt.Errorf("invalid id %d", id)
	}

	return id, nil
}
