package dbmodel

import "time"

type Todo struct {
	CreatedAt time.Time `bson:"createdAt"`
	Completed bool      `bson:"completed"`
	ID        string    `bson:"_id"`
	ListID    string    `bson:"listId"`
	Text      string    `bson:"text"`
}
