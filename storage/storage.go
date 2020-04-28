package storage

import (
	"github.com/autsakorn/go-shop/config"

	"go.mongodb.org/mongo-driver/mongo"
)
// Storage defines properties
type Storage struct {
	Product Product
}

// NewStorage connect mongo and return collection
func NewStorage(client *mongo.Client) (Storage, error) {

	env, _ := config.FromEnv()
	collection := client.Database(env.MongoDatabase).Collection(collectionName)

	return Storage{
		Product: ProductStorage{
			Collection: *collection,
		},
	}, nil
}