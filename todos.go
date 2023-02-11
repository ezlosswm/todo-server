package main

import (
	"time"
)

// creates a new todo item within a todo list
func NewTodo(item string, createdAt time.Time) *TodoList {
	return &TodoList{
		Item:      item,
		CreatedAt: createdAt,
	}
}

// todo
type TodoList struct {
	ID        int       `json:"id"`
	Item      string    `json:"item"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
}

type TodoReq struct {
	Item      string    `json:"item"`
	CreatedAt time.Time `json:"created_at"`
}
