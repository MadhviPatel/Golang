package models

import "time"

// TodoList schema of the todolist table
type TodoItemModel struct {
	Id          int
	Description string
	Created     time.Time
	Completed   bool
}
