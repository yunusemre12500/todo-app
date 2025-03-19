package conversion

import (
	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository/todorepository/mongodbtodorepository/internal/dbmodel"
)

func FromModel(todo *model.Todo) *dbmodel.Todo {
	return &dbmodel.Todo{
		CreatedAt: todo.CreatedAt,
		Completed: todo.Completed,
		ID:        string(todo.ID),
		ListID:    string(todo.ListID),
		Text:      todo.Text,
	}
}

func FromModels(todos []*model.Todo) []*dbmodel.Todo {
	var convertedTodos []*dbmodel.Todo

	for _, todo := range todos {
		convertedTodos = append(convertedTodos, FromModel(todo))
	}

	return convertedTodos
}

func ToModel(todo *dbmodel.Todo) *model.Todo {
	return &model.Todo{
		CreatedAt: todo.CreatedAt,
		Completed: todo.Completed,
		ID:        model.ID(todo.ID),
		ListID:    model.ID(todo.ListID),
		Text:      todo.Text,
	}
}

func ToModels(todos []*dbmodel.Todo) []*model.Todo {
	var convertedTodos []*model.Todo

	for _, todo := range todos {
		convertedTodos = append(convertedTodos, ToModel(todo))
	}

	return convertedTodos
}
