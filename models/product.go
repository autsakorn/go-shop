package models

import "time"

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Created time.Time `json:"created"`
}

type ProductStorage interface {
	Create(Product) error
	Update(Product) error
}