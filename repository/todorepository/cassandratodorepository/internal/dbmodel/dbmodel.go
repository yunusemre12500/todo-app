package dbmodel

import "time"

type Todo struct {
	Completed bool
	CreatedAt time.Time
	ID        string
	ListID    string
	Text      string
}
