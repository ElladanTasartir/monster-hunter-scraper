package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoClient(ctx context.Context, address, database string) (*Storage, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	return &Storage{
		client:   client,
		database: client.Database(database),
	}, nil
}

func (s *Storage) Disconnect(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
