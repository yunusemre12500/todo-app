package mongodbtodorepository

import (
	"context"

	"github.com/yunusemre12500/todo-app/model"
	"github.com/yunusemre12500/todo-app/repository"
	"github.com/yunusemre12500/todo-app/repository/todorepository"
	"github.com/yunusemre12500/todo-app/repository/todorepository/mongodbtodorepository/internal/conversion"
	"github.com/yunusemre12500/todo-app/repository/todorepository/mongodbtodorepository/internal/dbmodel"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/sync/singleflight"
)

type Repository struct {
	config                *Config
	getTodoByIDGroup      singleflight.Group
	getTodosByListIDGroup singleflight.Group
	todosCollection       *mongo.Collection
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

	r.todosCollection = client.Database(r.config.Database).Collection(r.config.Collection)

	return nil
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.todosCollection.Database().Client().Disconnect(ctx)
}

func (r *Repository) DeleteTodoByID(ctx context.Context, todoID model.ID) error {
	filter := bson.M{"_id": bson.M{"$eq": todoID}}

	if _, err := r.todosCollection.DeleteOne(ctx, filter); err != nil {
		if err == mongo.ErrNoDocuments {
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
	if _, err := r.todosCollection.InsertOne(ctx, conversion.FromModel(todo)); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTodoByID(ctx context.Context, todoID model.ID, todo *model.Todo) (*model.Todo, error) {
	filter := bson.M{"_id": bson.M{"$eq": todoID}}

	var updatedTodo dbmodel.Todo

	if err := r.todosCollection.FindOneAndUpdate(ctx, filter, todo).Decode(&updatedTodo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&updatedTodo), nil
}

func (r *Repository) getTodoByID(ctx context.Context, todoID model.ID) (*model.Todo, error) {
	filter := bson.M{"_id": bson.M{"$eq": todoID}}

	var todo dbmodel.Todo

	if err := r.todosCollection.FindOne(ctx, filter).Decode(&todo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return conversion.ToModel(&todo), nil
}

func (r *Repository) getTodosByListID(ctx context.Context, listID model.ID, opts *todorepository.GetTodosByListIDOptions) ([]*model.Todo, error) {
	filter := bson.M{"listId": bson.M{"$eq": listID}}

	var todos []*dbmodel.Todo

	findOpts := options.Find().
		SetLimit(int64(opts.Limit)).
		SetSkip(int64(opts.Offset))

	cursor, err := r.todosCollection.Find(ctx, filter, findOpts)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNoRecords
		}

		return nil, err

	}

	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return conversion.ToModels(todos), nil
}
