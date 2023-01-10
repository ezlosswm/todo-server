package main

import (
	"math/rand"
	"time"
)

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

func (t *TodoList) GetItemByID(id int) *TodoList {

	return nil
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
