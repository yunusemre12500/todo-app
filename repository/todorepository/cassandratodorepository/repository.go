package cassandratodorepository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
	"github.com/yunusemre12500/todo-app/repository/todorepository"
	"github.com/yunusemre12500/todo-app/repository/todorepository/cassandratodorepository/internal/conversion"
	"github.com/yunusemre12500/todo-app/repository/todorepository/cassandratodorepository/internal/dbmodel"
	"github.com/yunusemre12500/todo-app/repository/todorepository/cassandratodorepository/internal/queries"
	"golang.org/x/sync/singleflight"
)

type Repository struct {
	config                *Config
	getTodoByIDGroup      singleflight.Group
	getTodosByListIDGroup singleflight.Group
	session               *gocql.Session
}

func New(config *Config) *Repository {
	return &Repository{config: config}
}

func (r *Repository) Connect(_ context.Context) error {
	var err error

	r.session, err = r.config.toSession()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Disconnect(_ context.Context) error {
	r.session.Close()

	return nil
}

func (r *Repository) DeleteTodoByID(ctx context.Context, todoID model.ID) error {
	err := r.session.
		Query(fmt.Sprintf(queries.DeleteTodoByIDQuery, r.config.Keyspace, r.config.Table)).
		Bind(todoID).
		Exec()

	if err != nil {
		if err == gocql.ErrNotFound {
			return repository.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Repository) GetTodoByID(ctx context.Context, todoID model.ID) (*model.Todo, error) {
	todo, err, _ := r.getTodoByIDGroup.Do(string(todoID), func() (interface{}, error) {
		return r.getTodoByID(ctx, todoID)
	})

	return todo.(*model.Todo), err
}

func (r *Repository) GetTodosByListID(ctx context.Context, listID model.ID, opts *todorepository.GetTodosByListIDOptions) ([]*model.Todo, error) {
	todos, err, _ := r.getTodosByListIDGroup.Do(string(listID), func() (interface{}, error) {
		return r.getTodosByListID(ctx, listID, opts)
	})

	return todos.([]*model.Todo), err
}

func (r *Repository) SaveTodo(ctx context.Context, todo *model.Todo) error {
	err := r.session.
		Query(fmt.Sprintf(queries.SaveTodoQuery, r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(
			todo.Completed,
			todo.CreatedAt,
			todo.ID,
			todo.ListID,
			todo.Text,
		).
		Exec()

	if err != nil {
		if err, ok := err.(gocql.Error); ok && err.Code == gocql.ErrCodeAlreadyExists {
			return repository.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (r *Repository) UpdateTodoByID(ctx context.Context, todoID model.ID, todo *model.Todo) (*model.Todo, error) {
	err := r.session.
		Query(fmt.Sprintf(queries.UpdateTodoByIDQuery, r.config.Keyspace, r.config.Table)).
		Bind(todoID, todo.Completed, todo.CreatedAt, todo.ID, todo.ListID, todo.Text).
		Exec()

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return todo, nil
}

func (r *Repository) getTodoByID(ctx context.Context, todoID model.ID) (*model.Todo, error) {
	var todo dbmodel.Todo

	err := r.session.
		Query(fmt.Sprintf(queries.GetTodoByIDQuery, r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(todoID).
		Scan(
			&todo.Completed,
			&todo.CreatedAt,
			&todo.ID,
			&todo.ListID,
			&todo.Text,
		)

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&todo), nil
}

func (r *Repository) getTodosByListID(ctx context.Context, listID model.ID, opts *todorepository.GetTodosByListIDOptions) ([]*model.Todo, error) {
	var todo dbmodel.Todo

	iterator := r.session.
		Query(fmt.Sprintf(queries.GetTodosByListIDQuery, r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(listID, opts.Limit, opts.Offset).
		Iter()

	defer iterator.Close()

	scanner := iterator.Scanner()

	var todos []*dbmodel.Todo

	for scanner.Next() {
		if err := scanner.Scan(
			todo.Completed,
			todo.CreatedAt,
			todo.ID,
			todo.ListID,
			todo.Text,
		); err != nil {
			return nil, err
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	if len(todos) == 0 {
		return nil, repository.ErrNoRecords
	}

	return conversion.ToModels(todos), nil
}
