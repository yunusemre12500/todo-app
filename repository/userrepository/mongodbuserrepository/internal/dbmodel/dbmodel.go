package dbmodel

import "time"

type User struct {
	CreatedAt    time.Time `bson:"createdAt"`
	DisplayName  string    `bson:"displayName"`
	EmailAddress string    `bson:"emailAddress"`
	ID           string    `bson:"_id"`
	Name         string    `bson:"name"`
	PasswordHash string    `bson:"passwordHash"`
}
