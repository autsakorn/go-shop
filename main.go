// main.go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/autsakorn/go-shop/controllers"
	"github.com/autsakorn/go-shop/middleware"
	myMongo "github.com/autsakorn/go-shop/mongo"
	"github.com/autsakorn/go-shop/services"
	"github.com/autsakorn/go-shop/storage"
	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func routes(c controllers.Controllers) {
	router = httprouter.New()

	router.GET("/health-check", c.HealthCheck.GetHealthCheck)
	router.POST("/product", c.Product.CreateProduct)
	router.DELETE("/product", c.Product.DeleteProduct)
	router.GET("/product", c.Product.FindProducts)
	router.GET("/product/:id", c.Product.FindProductByID)
	router.PATCH("/product", c.Product.UpdateProduct)
	router.PUT("/product", c.Product.UpsertProduct)
}

func main() {
	// Connect mongo client
	ms, _ := myMongo.NewMongoStorage(context.Background())
	// Create Storage
	Storage, _ := storage.NewStorage(ms.Client)
	// Create ProductService
	ProductService := services.NewProductService()

	c := &controllers.Controllers{
		Product:     controllers.Product{ProductService, Storage},
		HealthCheck: controllers.HealthCheck{},
	}
	routes(*c)
	m := middleware.NewContentTypeMiddleware(router)
	log.Fatal(http.ListenAndServe(":8011", m))
}
