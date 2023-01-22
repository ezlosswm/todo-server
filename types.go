package main

import "time"

var ListOfTodo TodoList

type TodoItem struct {
	ID          int       `json:"id"`
	Activity    string    `json:"activity"`
	CompletedAt time.Time `json:"completed_at"`
	Completed   bool      `json:"completed"`
}

type TodoList struct {
	Items []TodoItem `json:"items"`
}
