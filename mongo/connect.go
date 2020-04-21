package mongo

import (
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/autsakorn/go-shop/config"
)

type MongoStorage struct {
	*mongo.Client
}

func NewMongoStorage(ctx context.Context) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	cfg, _ := config.FromEnv()

	defer cancel()
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		cfg.MongoUsername,
		cfg.MongoPassword,
		cfg.MongoHost,
		cfg.MongoPort,
	)
	
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	ms := MongoStorage{
		Client: client,
	}
	return &ms, nil
}
