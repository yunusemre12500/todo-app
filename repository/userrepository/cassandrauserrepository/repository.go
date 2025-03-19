package cassandrauserrepository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
	"github.com/yunusemre12500/todo-app/repository/userrepository/cassandrauserrepository/internal/conversion"
	"github.com/yunusemre12500/todo-app/repository/userrepository/cassandrauserrepository/internal/dbmodel"
	"github.com/yunusemre12500/todo-app/repository/userrepository/cassandrauserrepository/internal/queries"
	"golang.org/x/sync/singleflight"
)

type Repository struct {
	config             *Config
	getUserByIDGroup   singleflight.Group
	getUserByNameGroup singleflight.Group
	session            *gocql.Session
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

func (r *Repository) DeleteUserByID(ctx context.Context, userID model.ID) error {
	err := r.session.
		Query(fmt.Sprintf(r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(userID).
		Exec()

	if err != nil {
		if err == gocql.ErrNotFound {
			return repository.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, userID model.ID) (*model.User, error) {
	user, err, _ := r.getUserByIDGroup.Do(string(userID), func() (interface{}, error) {
		return r.getUserByID(ctx, userID)
	})

	return user.(*model.User), err
}

func (r *Repository) GetUserByName(ctx context.Context, userName string) (*model.User, error) {
	user, err, _ := r.getUserByNameGroup.Do(userName, func() (interface{}, error) {
		return r.getUserByName(ctx, userName)
	})

	return user.(*model.User), err
}

func (r *Repository) SaveUser(ctx context.Context, user *model.User) error {
	println(queries.SaveUserQuery)
	err := r.session.
		Query(fmt.Sprintf(queries.SaveUserQuery, r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(
			user.CreatedAt,
			user.DisplayName,
			user.EmailAddress,
			user.ID,
			user.Name,
			user.PasswordHash,
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

func (r *Repository) getUserByID(ctx context.Context, userID model.ID) (*model.User, error) {
	var user dbmodel.User

	err := r.session.
		Query(fmt.Sprintf(queries.GetUserByIDQuery, r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(userID).
		Scan(&user.CreatedAt, &user.DisplayName, &user.EmailAddress, &user.ID, &user.Name, &user.PasswordHash)

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&user), nil
}

func (r *Repository) getUserByName(ctx context.Context, userName string) (*model.User, error) {
	var user dbmodel.User

	err := r.session.
		Query(fmt.Sprintf(queries.GetUserByNameQuery, r.config.Keyspace, r.config.Table)).
		WithContext(ctx).
		Bind(userName).
		Scan(&user.CreatedAt, &user.DisplayName, &user.EmailAddress, &user.ID, &user.Name, &user.PasswordHash)

	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&user), nil
}
