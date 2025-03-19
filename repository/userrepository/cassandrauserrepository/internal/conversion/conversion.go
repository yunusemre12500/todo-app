package conversion

import (
	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository/userrepository/cassandrauserrepository/internal/dbmodel"
)

func FromModel(user *model.User) *dbmodel.User {
	return &dbmodel.User{
		CreatedAt:    user.CreatedAt,
		DisplayName:  user.DisplayName,
		EmailAddress: user.EmailAddress,
		ID:           string(user.ID),
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
	}
}

func ToModel(user *dbmodel.User) *model.User {
	return &model.User{
		CreatedAt:    user.CreatedAt,
		DisplayName:  user.DisplayName,
		EmailAddress: user.EmailAddress,
		ID:           model.ID(user.ID),
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
	}
}
