package conversion

import (
	"github.com/yunusemre12500/todo-app/model"
	dbmodel "github.com/yunusemre12500/todo-app/repository/user/mongo/internal/model"
)

func FromDomain(user *model.User) *dbmodel.User {
	return &dbmodel.User{
		CreatedAt:    user.CreatedAt,
		DisplayName:  user.DisplayName,
		EmailAddress: user.EmailAddress,
		ID:           user.ID,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
	}
}

func ToDomain(user *dbmodel.User) *model.User {
	return &model.User{
		CreatedAt:    user.CreatedAt,
		DisplayName:  user.DisplayName,
		EmailAddress: user.EmailAddress,
		ID:           user.ID,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
	}
}
