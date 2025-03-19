package model

import "time"

type User struct {
	CreatedAt    time.Time
	DisplayName  string
	EmailAddress string
	ID           ID
	Name         string
	PasswordHash string
}
