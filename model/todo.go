package model

import "time"

type Todo struct {
	Completed bool
	CreatedAt time.Time
	ID        ID
	ListID    ID
	Text      string
}
