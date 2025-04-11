package repository

import "context"

type Repository interface {
	Connect(context.Context) error
	Disconnect(context.Context) error
}
