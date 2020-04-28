package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/autsakorn/go-shop/services"
	"github.com/autsakorn/go-shop/storage"
	"github.com/autsakorn/go-shop/types"
	"github.com/julienschmidt/httprouter"
)

// Product defines properties controller
type Product struct {
	ProductService services.Product
	Storage        storage.Storage
}

// CreateProduct create a new product
func (p Product) CreateProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var newProduct types.InputCreateProduct
	decoder.Decode(&newProduct)

	// Call service
	results, _ := p.ProductService.CreateProduct(newProduct, p.Storage.Product)
	json.NewEncoder(w).Encode(results)
}

// DeleteProduct delete a product
func (p Product) DeleteProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var deleteProduct types.InputDeleteProduct
	decoder.Decode(&deleteProduct)

	// Call service
	results, _ := p.ProductService.DeleteProduct(deleteProduct, p.Storage.Product)
	json.NewEncoder(w).Encode(results)
}

// FindProducts find a product with paging
func (p Product) FindProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	// Call service
	results, _ := p.ProductService.FindProducts(page, limit, p.Storage.Product)
	json.NewEncoder(w).Encode(results)
}

// FindProductByID find a product by ID
func (p Product) FindProductByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	results, _ := p.ProductService.FindProductByID(id, p.Storage.Product)
	json.NewEncoder(w).Encode(results)
}

// UpdateProduct update a product by ID
func (p Product) UpdateProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var updateProduct types.InputProduct
	decoder.Decode(&updateProduct)

	// Call service
	results, _ := p.ProductService.UpdateProduct(updateProduct, p.Storage.Product)
	json.NewEncoder(w).Encode(results)
}

// UpsertProduct create or update if exists
func (p Product) UpsertProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var updateProduct types.InputProduct
	decoder.Decode(&updateProduct)

	// Call service
	results, _ := p.ProductService.UpsertProduct(updateProduct, p.Storage.Product)
	json.NewEncoder(w).Encode(results)
}
