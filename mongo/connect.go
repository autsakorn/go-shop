package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/autsakorn/go-shop/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage of mongo client and collection
type Storage struct {
	*mongo.Client
}

// NewMongoStorage connect db and collection
func NewMongoStorage(ctx context.Context) (*Storage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	env, _ := config.FromEnv()

	defer cancel()
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		env.MongoUsername,
		env.MongoPassword,
		env.MongoHost,
		env.MongoPort,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	ms := Storage{
		Client: client,
	}
	return &ms, nil
}
