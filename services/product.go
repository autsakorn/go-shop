package services

import (
	"errors"
	"fmt"

	"github.com/autsakorn/go-shop/storage"
	"github.com/autsakorn/go-shop/types"
)

// CalSkip function for calulate skip
func CalSkip(page int64, limit int64) int64 {
	if page > 0 {
		return (page - 1) * limit
	}
	return 0
}

// CreateProduct product service
func CreateProduct(
	product types.InputCreateProduct,
	sProduct storage.Product,
) (types.OutputCreateProduct, error) {
	var results types.OutputCreateProduct
	results.Created = false
	// Validate
	if product.Price < 1 {
		priceIsRequiredMessage := "Price is required"
		results.Message = fmt.Sprintf(priceIsRequiredMessage)
		return results, errors.New(priceIsRequiredMessage)
	}

	err := sProduct.Create(product)

	if err != nil {
		results.Message = fmt.Sprintf("Product %v", err)
		return results, err
	}
	results.Created = true
	return results, nil
}

// DeleteProduct delete product service
func DeleteProduct(
	product types.InputDeleteProduct,
	sProduct storage.Product,
) (types.OutputDeleteProduct, error) {
	var results types.OutputDeleteProduct
	results.Deleted = false
	// Validate
	if product.ID == "" {
		idIsRequiredMessage := "ID is required"
		results.Message = fmt.Sprintf(idIsRequiredMessage)
		return results, errors.New(idIsRequiredMessage)
	}
	err := sProduct.Delete(product)
	if err != nil {
		results.Message = fmt.Sprintf("Product %v", err)
		return results, err
	}
	results.Deleted = true
	return results, nil
}

// FindProductByID find product by ID
func FindProductByID(
	id string,
	sProduct storage.Product,
) (types.OutputProduct, error) {

	var results types.OutputProduct
	// Validate
	if id == "" {
		idIsRequiredMessage := "ID is required"
		results.Message = fmt.Sprintf(idIsRequiredMessage)
		return results, errors.New(idIsRequiredMessage)
	}
	product, err := sProduct.FindByID(id)
	if err != nil {
		results.Message = fmt.Sprintf("Product %v", err)
		return results, err
	}
	results.Data = product
	return results, nil
}

// FindProducts find all product
func FindProducts(
	page int64,
	limit int64,
	sProduct storage.Product,
) (types.OutputProducts, error) {
	var results types.OutputProducts
	// Set Default current page
	if page < 1 {
		page = 1
	}
	// Set Default limit per page
	if limit < 1 {
		limit = 10
	}
	skip := CalSkip(page, limit)
	products, err := sProduct.Find(skip, limit)
	if err != nil {
		return results, err
	}
	var input types.InputProduct
	totals, _ := sProduct.Count(input)
	results.Totals = totals

	results.Data = products
	return results, nil
}

// UpdateProduct update product by ID
func UpdateProduct(
	product types.InputProduct,
	sProduct storage.Product,
) (types.OutputUpdateProduct, error) {

	var results types.OutputUpdateProduct
	err := sProduct.Update(product)
	results.Created = false
	results.Updated = false
	if err != nil {
		results.Message = fmt.Sprintf("Product %v", err)
		return results, err
	}
	results.Updated = true
	return results, nil
}

// UpsertProduct service, create or update if exists
func UpsertProduct(
	product types.InputProduct,
	sProduct storage.Product,
) (types.OutputUpdateProduct, error) {
	var results types.OutputUpdateProduct
	results.Created = false
	results.Updated = false
	numberProduct, err := sProduct.Count(product)
	if err != nil {
		results.Message = fmt.Sprintf("Product %v", err)
		return results, err
	}
	if numberProduct < 1 {
		newProduct := types.InputCreateProduct{
			Name:  product.Name,
			Price: product.Price,
		}
		if err := sProduct.Create(newProduct); err != nil {
			results.Message = fmt.Sprintf("Product %v", err)
			return results, err
		}
		results.Created = true
	} else {
		if err := sProduct.Update(product); err != nil {
			results.Message = fmt.Sprintf("Product %v", err)
			return results, err
		}
		results.Updated = true
	}
	return results, nil
}
