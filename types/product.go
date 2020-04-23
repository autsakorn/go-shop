package types

import (
	"github.com/autsakorn/go-shop/models"
)

// InputCreateProduct data for create new product
type InputCreateProduct struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// InputDeleteProduct data for delete product
type InputDeleteProduct struct {
	ID string `json:"id"`
}

// InputProduct data for update and upsert product
type InputProduct struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// OutputCreateProduct response data create product
type OutputCreateProduct struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
}

// OutputDeleteProduct response data delete product
type OutputDeleteProduct struct {
	Deleted bool   `json:"deleted"`
	Message string `json:"message"`
}

// OutputUpdateProduct response data update product
type OutputUpdateProduct struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
	Updated bool   `json:"updated"`
}

// OutputProduct response data find product
type OutputProduct struct {
	Message string            `json:"message"`
	Data    []*models.Product `json:"data"`
}

// OutputProducts response data find product
type OutputProducts struct {
	Message string            `json:"message"`
	Totals  int64             `json:"totals"`
	Data    []*models.Product `json:"data"`
}
