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
	*mongo.Collection
}

// NewMongoStorage connect db and collection
func NewMongoStorage(ctx context.Context, collection string) (*Storage, error) {
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

	ms := Storage{
		Client:     client,
		Collection: client.Database(cfg.MongoDatabase).Collection(collection),
	}
	return &ms, nil
}
