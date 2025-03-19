package mongodbuserrepository

import (
	"context"

	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
	"github.com/yunusemre12500/todo-app/repository/userrepository/mongodbuserrepository/internal/conversion"
	"github.com/yunusemre12500/todo-app/repository/userrepository/mongodbuserrepository/internal/dbmodel"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/sync/singleflight"
)

type Repository struct {
	config             *Config
	getUserByIDGroup   singleflight.Group
	getUserByNameGroup singleflight.Group
	usersCollection    *mongo.Collection
}

func New(config *Config) *Repository {
	return &Repository{config: config}
}

func (r *Repository) Connect(ctx context.Context) error {
	client, err := r.config.toClient()

	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	r.usersCollection = client.Database(r.config.Database).Collection(r.config.Collection)

	return nil
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.usersCollection.Database().Client().Disconnect(ctx)
}

func (r *Repository) DeleteUserByID(ctx context.Context, userID model.ID) error {
	filter := bson.M{"_id": bson.M{"$eq": userID}}

	if _, err := r.usersCollection.DeleteOne(ctx, filter); err != nil {
		if err == mongo.ErrNoDocuments {
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
	if _, err := r.usersCollection.InsertOne(ctx, conversion.FromModel(user)); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return repository.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (r *Repository) getUserByID(ctx context.Context, userID model.ID) (*model.User, error) {
	filter := bson.M{"_id": bson.M{"$eq": userID}}

	var user dbmodel.User

	if err := r.usersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&user), nil
}

func (r *Repository) getUserByName(ctx context.Context, userName string) (*model.User, error) {
	filter := bson.M{"name": bson.M{"$eq": userName}}

	var user dbmodel.User

	if err := r.usersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&user), nil
}
