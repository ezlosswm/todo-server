package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

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

func (t *TodoList) GetItemByID(id int) (*TodoItem, error) {
	for itemID := range t.Items {
		if id == t.Items[itemID].ID {
			return &t.Items[itemID], nil
		}
	}

	return nil, fmt.Errorf("item with id %d not found", id)
}

func (t *TodoList) AddItem(item string) *TodoList {
	todoList := NewTodo(item)
	t.Items = append(t.Items, todoList.Items...)

	return todoList
}

func (t *TodoList) UpdateItem(id int, newItem TodoItem) *TodoList {

	for i, todo := range t.Items {
		if todo.ID == id {

			todo.ID = newItem.ID
			todo.Activity = newItem.Activity
			todo.CompletedAt = newItem.CompletedAt
			todo.Completed = newItem.Completed

			t.Items[i] = newItem
			break
		}
	}

	return t
}

func (t *TodoList) CompleteItem(id int, completedItem []TodoItem) *TodoList {
	for i, todo := range t.Items {
		if todo.ID == id {
			completedItem[i].CompletedAt = time.Now()
			completedItem[i].Completed = true

			todo.CompletedAt = completedItem[i].CompletedAt
			todo.Completed = completedItem[i].Completed
			break
		}
	}

	return t
}

func (t *TodoList) DeleteItem(id int) error {
	for i, todo := range t.Items {
		if todo.ID == id {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("id %s not found", strconv.Itoa(id))
}

func GetID(r *http.Request) (int, error) {
	// retrieves the id from the post data
	idFromUser := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idFromUser)
	if err != nil {
		return id, fmt.Errorf("%d not a valid id", id)
	}

	return id, nil
}
