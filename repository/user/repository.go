package user

import (
	"context"

	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
)

type Repository interface {
	repository.Repository
	DeleteByID(context.Context, model.ID) error
	GetByID(context.Context, model.ID) (*model.User, error)
	Save(context.Context, *model.User) error
	UpdateByID(context.Context, model.ID, *UpdateByIDOptions) (*model.User, error)
}

type UpdateByIDOptions struct {
	DisplayName  string
	EmailAddress string
	Name         string
	PasswordHash string
}
