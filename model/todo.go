package model

import "time"

type Todo struct {
	CreatedAt time.Time
	Completed bool
	ID        ID
	ListID    ID
	Text      string
}
