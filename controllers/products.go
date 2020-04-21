package controllers

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/autsakorn/go-shop/services"
	"github.com/autsakorn/go-shop/models"
	"github.com/autsakorn/go-shop/storage"
)

func GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, _ := services.GetProducts()
	json.NewEncoder(w).Encode(results)
}

func GetProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	log.Print(id)
	results, _ := services.GetProduct()
	json.NewEncoder(w).Encode(results)
}

func PostProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var newProduct models.Product
	decoder.Decode(&newProduct)

	// Call Storage [storage/product]
	productStrg, _ := storage.NewProductStorage()
	
	// Call service
	results, _ := services.PostProduct(newProduct, productStrg)
	json.NewEncoder(w).Encode(results)
}

func PatchProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(req.Body)
	var patchProduct models.Product
	decoder.Decode(&patchProduct)

	// Call service
	results, _ := services.PatchProduct(patchProduct)
	json.NewEncoder(w).Encode(results)
}