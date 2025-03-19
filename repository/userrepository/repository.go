package userrepository

import (
	"context"

	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
)

type Repository interface {
	repository.Repository
	DeleteUserByID(context.Context, model.ID) error
	GetUserByID(context.Context, model.ID) (*model.User, error)
	GetUserByName(context.Context, string) (*model.User, error)
	SaveUser(context.Context, *model.User) error
}
