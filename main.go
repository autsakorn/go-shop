// main.go
package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/autsakorn/go-shop/controllers"
	"github.com/autsakorn/go-shop/middleware"
)

var router *httprouter.Router


func routes() {
	router = httprouter.New()
	router.GET("/product", controllers.GetProducts)
	router.GET("/product/:id", controllers.GetProduct)
	router.POST("/product", controllers.PostProduct)
	router.PATCH("/product", controllers.PatchProduct)
}

func main() {
	routes()
	m := middleware.NewContentTypeMiddleware(router)
	log.Fatal(http.ListenAndServe(":8000", m))
}