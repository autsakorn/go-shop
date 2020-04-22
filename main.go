// main.go
package main

import (
	"log"
	"net/http"

	"github.com/autsakorn/go-shop/controllers"
	"github.com/autsakorn/go-shop/middleware"
	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func routes() {
	router = httprouter.New()
	router.POST("/product", controllers.CreateProduct)
	router.DELETE("/product", controllers.DeleteProduct)
	router.GET("/product", controllers.FindProducts)
	router.GET("/product/:id", controllers.FindProductByID)
	router.PATCH("/product", controllers.UpdateProduct)
	router.PUT("/product", controllers.UpsertProduct)
}

func main() {
	routes()
	m := middleware.NewContentTypeMiddleware(router)
	log.Fatal(http.ListenAndServe(":8000", m))
}
