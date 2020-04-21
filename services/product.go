package services

import (
	"log"
	"time"
	"github.com/autsakorn/go-shop/models"
)

func GetProduct() ([]*models.Product, error) {

	var results []*models.Product	
	var result = models.Product{
		ID: 1,
		Name: "Apple",
	}
	results = append(results, &result)
	return results, nil 
}

func GetProducts() ([]*models.Product, error) {

	results := []*models.Product{
		{
			ID: 1,
			Name: "Apple",
		},
		{
			ID: 2,
			Name: "Google",
		},
	}
	return results, nil
}

func PostProduct(product models.Product, productStrg models.ProductStorage) ([]*models.Product, error) {
	var results []*models.Product
	product.Created = time.Now()
	log.Print("Service", product)

	err := productStrg.Create(product)
	
	if err != nil {
		return []*models.Product{}, err
	}
	results = []*models.Product{
		{
			ID: product.ID,
			Name: product.Name,
			Created: product.Created,
		},
	}
	return results, nil
}

func PatchProduct(product models.Product) ([]*models.Product, error) {
	var results []*models.Product
	product.Created = time.Now()
	log.Print(product)
	results = []*models.Product{
		{
			ID: product.ID,
			Name: product.Name,
			Created: product.Created,
		},
	}
	return results, nil
}