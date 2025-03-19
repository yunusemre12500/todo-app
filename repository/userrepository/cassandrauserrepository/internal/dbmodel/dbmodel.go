package dbmodel

import "time"

type User struct {
	CreatedAt    time.Time
	DisplayName  string
	EmailAddress string
	ID           string
	Name         string
	PasswordHash string
}
