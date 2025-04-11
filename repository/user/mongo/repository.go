package mongo

import (
	"context"

	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
	userrepository "github.com/yunusemre12500/todo-app/repository/user"
	"github.com/yunusemre12500/todo-app/repository/user/mongo/internal/conversion"
	dbmodel "github.com/yunusemre12500/todo-app/repository/user/mongo/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ userrepository.Repository = (*Repository)(nil)

type Repository struct {
	collection *mongo.Collection
	config     *Config
}

func New(config *Config) *Repository {
	return &Repository{
		config: config,
	}
}

func (r *Repository) Connect(ctx context.Context) error {
	client, err := r.config.toClient()

	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	r.collection = client.Database(r.config.Database).Collection(r.config.Collection)

	return nil
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.collection.Database().Client().Disconnect(ctx)
}

func (r *Repository) DeleteByID(ctx context.Context, userID model.ID) error {
	filter := bson.M{"_id": userID}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		if err == mongo.ErrNoDocuments {
			return repository.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, userID model.ID) (*model.User, error) {
	filter := bson.M{"_id": userID}

	var user dbmodel.User

	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToDomain(&user), nil
}

func (r *Repository) Save(ctx context.Context, user *model.User) error {
	if _, err := r.collection.InsertOne(ctx, conversion.FromDomain(user)); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return repository.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (r *Repository) UpdateByID(ctx context.Context, userID model.ID, opts *userrepository.UpdateByIDOptions) (*model.User, error) {
	filter := bson.M{"_id": userID}

	update := bson.M{"$set": bson.M{
		"displayName":  opts.DisplayName,
		"emailAddress": opts.EmailAddress,
		"name":         opts.Name,
		"passwordHash": opts.PasswordHash,
	}}

	var updatedUser dbmodel.User

	if err := r.collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		} else if mongo.IsDuplicateKeyError(err) {
			return nil, repository.ErrAlreadyExists
		}

		return nil, err
	}

	return conversion.ToDomain(&updatedUser), nil
}
