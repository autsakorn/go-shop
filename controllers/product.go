package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	myMongo "github.com/autsakorn/go-shop/mongo"
	"github.com/autsakorn/go-shop/services"
	"github.com/autsakorn/go-shop/storage"
	"github.com/autsakorn/go-shop/types"
	"github.com/julienschmidt/httprouter"
)

// CreateProduct create a new product
func CreateProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var newProduct types.InputCreateProduct
	decoder.Decode(&newProduct)

	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())

	// Call Storage [storage/product]
	sProduct, _ := storage.NewProductStorage(ms.Client)

	// Call service
	results, _ := services.CreateProduct(newProduct, sProduct)
	json.NewEncoder(w).Encode(results)
}

// DeleteProduct delete a product
func DeleteProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var deleteProduct types.InputDeleteProduct
	decoder.Decode(&deleteProduct)

	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())

	// Call Storage [storage/product]
	sProduct, _ := storage.NewProductStorage(ms.Client)

	// Call service
	results, _ := services.DeleteProduct(deleteProduct, sProduct)
	json.NewEncoder(w).Encode(results)
}

// FindProducts find a product with paging
func FindProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())

	// Call Storage [storage/product]
	sProduct, _ := storage.NewProductStorage(ms.Client)
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	// Call service
	results, _ := services.FindProducts(page, limit, sProduct)
	json.NewEncoder(w).Encode(results)
}

// FindProductByID find a product by ID
func FindProductByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())

	// Call Storage [storage/product]
	sProduct, _ := storage.NewProductStorage(ms.Client)

	results, _ := services.FindProductByID(id, sProduct)
	json.NewEncoder(w).Encode(results)
}

// UpdateProduct update a product by ID
func UpdateProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var updateProduct types.InputProduct
	decoder.Decode(&updateProduct)

	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())

	// Call Storage [storage/product]
	sProduct, _ := storage.NewProductStorage(ms.Client)

	// Call service
	results, _ := services.UpdateProduct(updateProduct, sProduct)
	json.NewEncoder(w).Encode(results)
}

// UpsertProduct create or update if exists
func UpsertProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var updateProduct types.InputProduct
	decoder.Decode(&updateProduct)

	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())

	// Call Storage [storage/product]
	sProduct, _ := storage.NewProductStorage(ms.Client)

	// Call service
	results, _ := services.UpsertProduct(updateProduct, sProduct)
	json.NewEncoder(w).Encode(results)
}
