package model

import (
	"time"

	"github.com/yunusemre12500/todo-app/model"
)

type User struct {
	CreatedAt    time.Time `bson:"createdAt"`
	DisplayName  string    `bson:"displayName"`
	EmailAddress string    `bson:"emailAddress"`
	ID           model.ID  `bson:"_id"`
	Name         string    `bson:"name"`
	PasswordHash string    `bson:"passwordHash"`
}
