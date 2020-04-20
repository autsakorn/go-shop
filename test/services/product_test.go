package product_test

import (
	"testing"
	"time"
	"log"
	"github.com/autsakorn/go-shop/services"
	"github.com/autsakorn/go-shop/models"
)

var insertProductMock func(product models.Product) error
var updateProductMock func(product models.Product) error

type productStorageMock struct{}

func (productStrgMock productStorageMock) InsertProduct(product models.Product) error {
	return insertProductMock(product)
}
func (productStrgMock productStorageMock) UpdateProduct(product models.Product) error {
	return updateProductMock(product)
}

func TestPostProduct(t *testing.T) {
	tests := []struct {
		name string
		data models.Product
		want error
	}{
		{
			"base-case",
			models.Product{
				ID: 1,
				Name: "Mas",
				Created: time.Now(),
			},
			nil,
		},
	}
	insertProductMock = func(product models.Product) error {
		return nil
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		_, err := services.PostProduct(test.data, productStorage)
		log.Print(err)
		if err != test.want {
			t.Errorf("PostProduct Fail = %v, want %v", err, test.want)
		}
	}
}