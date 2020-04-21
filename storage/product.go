package storage

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/autsakorn/go-shop/models"
	myMongo "github.com/autsakorn/go-shop/mongo"
)

type productStorage struct {
	mongo.Client
	DB string
	Collection string
}

func NewProductStorage() (models.ProductStorage, error) {
	db := "qwertyuiop"
	collection := "products"
	ms, err := myMongo.NewMongoStorage(context.Background())
	if err != nil {
		return nil, err
	}
	return productStorage{
		*ms.Client,
		db,
		collection,
	}, nil
}

func (ps productStorage) Create(product models.Product) error {
	coll := ps.Client.Database(ps.DB).Collection(ps.Collection)
	if _, err := coll.InsertOne(context.Background(), product); err != nil {
		log.Print("err", err)
		return err
	}
	return nil
}

func (p productStorage) Update(product models.Product) error {
	log.Print("Storage", product)
	return nil
}
