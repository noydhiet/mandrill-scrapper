package storage

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type StorageMongo struct {
	client *mongo.Client
}

// NewStorageMongo creates a new instance of StorageMongo
// dsn "mongodb://localhost:27017"
func NewStorageMongo(dsn string) (Storage, error) {
	dsnDefault := "mongodb://localhost:27017"
	if dsn != "" {
		dsnDefault = dsn
	}

	client, err := mongo.Connect(options.Client().ApplyURI(dsnDefault))
	if err != nil {
		return nil, err
	}
	return &StorageMongo{
		client: client,
	}, nil
}

func (s *StorageMongo) Store(ctx context.Context, model string, data interface{}) error {
	collection := s.client.Database("mandrill").Collection(model)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return errors.Wrap(err, "failed to insert data")
	}
	return nil
}
func (s *StorageMongo) Find(ctx context.Context, model string, filter interface{}, data interface{}) error {
	collection := s.client.Database("mandrill").Collection(model)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "failed to find data")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, data); err != nil {
		return errors.Wrap(err, "failed to decode data")
	}
	return nil
}
