package services

import (
	"testing"
	"time"
	"log"
	"github.com/autsakorn/go-shop/models"
)

var createMock func(product models.Product) error
var updateMock func(product models.Product) error

type productStorageMock struct{}

func (productStrgMock productStorageMock) Create(product models.Product) error {
	return createMock(product)
}
func (productStrgMock productStorageMock) Update(product models.Product) error {
	return updateMock(product)
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
	createMock = func(product models.Product) error {
		return nil
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		_, err := PostProduct(test.data, productStorage)
		log.Print(err)
		if err != test.want {
			t.Errorf("PostProduct Fail = %v, want %v", err, test.want)
		}
	}
}