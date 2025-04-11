package model

import "time"

type TodoList struct {
	CreatedAt time.Time
	ID        ID
	Name      string
	UserID    ID
}
