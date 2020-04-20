package storage

import (
	"log"
	"github.com/autsakorn/go-shop/models"
)

type productStorage struct {}

func NewProductStorage() models.ProductStorage {
	return productStorage{}
}

func (p productStorage) InsertProduct(product models.Product) error {
	log.Print("Storage", product)
	return nil
}

func (p productStorage) UpdateProduct(product models.Product) error {
	log.Print("Storage", product)
	return nil
}
