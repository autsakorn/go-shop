package models

import "time"

type Product struct {
	ID				int 				`json:"id"`
	Name 			string			`json:"name"`
	Created 	time.Time 	`json:"created"`
}

type ProductStorage interface {
	InsertProduct(Product) error
	UpdateProduct(Product) error
}