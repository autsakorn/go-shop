package storage

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/autsakorn/go-shop/models"
	myMongo "github.com/autsakorn/go-shop/mongo"
	"github.com/autsakorn/go-shop/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "products"

// Product represents all possible actions
type Product interface {
	Create(types.InputCreateProduct) error
	Count(types.InputProduct) (int64, error)
	Delete(types.InputDeleteProduct) error
	FindByID(string) ([]*models.Product, error)
	Find(int64, int64) ([]*models.Product, error)
	Update(types.InputProduct) error
}

type storage struct {
	mongo.Collection
}

// NewProductStorage connect mongo and return collection
func NewProductStorage() (Product, error) {
	ms, err := myMongo.NewMongoStorage(context.Background(), collectionName)
	if err != nil {
		return nil, err
	}
	return storage{*ms.Collection}, nil
}

// Create new product
func (s storage) Create(product types.InputCreateProduct) error {
	// Prepare data create
	create := models.Product{
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Execute create
	if _, err := s.Collection.InsertOne(context.Background(), create); err != nil {
		log.Printf("Product Create Error %v", err)
		return err
	}
	return nil
}

// Count products
func (s storage) Count(
	product types.InputProduct,
) (int64, error) {
	filter := bson.M{}
	if product.ID != "" {
		objectID, _ := primitive.ObjectIDFromHex(product.ID)
		filter["_id"] = objectID
	}

	numberRows, err := s.Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Printf("Product Count Error %v", err)
		return 0, err
	}
	return numberRows, nil
}

// Delete delete
func (s storage) Delete(product types.InputDeleteProduct) error {
	// Prepare filter _id
	objectID, _ := primitive.ObjectIDFromHex(product.ID)
	filter := bson.M{"_id": objectID}
	// Execute delete
	deleteResult, err := s.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Product Delete Error %v", err)
		return err
	}
	if deleteResult.DeletedCount < 1 {
		return errors.New("Not found")
	}
	return nil
}

// Find all product set based on given skip and limit
func (s storage) Find(skip int64, limit int64) ([]*models.Product, error) {
	ctx := context.Background()
	var products []*models.Product
	// Set filter and find option
	filter := bson.M{}
	findOption := options.Find()
	findOption.SetLimit(limit)
	findOption.SetSkip(skip)
	cur, err := s.Collection.Find(ctx, filter, findOption)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.Product
		err := cur.Decode(&result)
		if err != nil {
			return products, err
		}
		products = append(products, &result)
	}
	if err := cur.Err(); err != nil {
		log.Printf("Product Find Error %v", err)
		return nil, err
	}
	return products, nil
}

// FindByID find one product by ID
func (s storage) FindByID(id string) ([]*models.Product, error) {
	var products []*models.Product
	var result models.Product

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	cur := s.Collection.FindOne(context.Background(), filter)
	err := cur.Decode(&result)
	if err != nil {
		log.Printf("Product FindByID Error %v", err)
		return products, err
	}
	products = append(products, &result)
	return products, nil
}

// Update product by _id
func (s storage) Update(product types.InputProduct) error {

	objectID, _ := primitive.ObjectIDFromHex(product.ID)
	filter := bson.M{"_id": objectID}
	updateField := bson.M{}
	updateField["price"] = product.Price
	updateField["name"] = product.Name
	updateField["updated_at"] = time.Now()
	update := bson.M{"$set": updateField}

	// Execute update
	updateResult, err := s.Collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Printf("Product update error %v", err)
		return err
	}

	if updateResult.MatchedCount < 1 {
		return errors.New("Not found")
	}
	return nil
}
