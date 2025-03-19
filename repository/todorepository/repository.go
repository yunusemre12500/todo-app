package todorepository

import (
	"context"

	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
)

type Repository interface {
	repository.Repository
	DeleteTodoByID(context.Context, model.ID) error
	GetTodoByID(context.Context, model.ID) (*model.ID, error)
	GetTodosByListID(context.Context, model.ID, *GetTodosByListIDOptions) ([]*model.Todo, error)
	SaveTodo(context.Context, *model.Todo) error
	UpdateTodoByID(context.Context, model.ID, *model.Todo) (*model.Todo, error)
}

type GetTodosByListIDOptions struct {
	Limit  uint64
	Offset uint64
}
