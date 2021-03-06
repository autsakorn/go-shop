package types

import (
	"github.com/autsakorn/go-shop/models"
)

// InputCreateProduct defines properties of input create new product
type InputCreateProduct struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// InputDeleteProduct defines properties of input delete product
type InputDeleteProduct struct {
	ID string `json:"id"`
}

// InputProduct defines properties of input update and upsert product
type InputProduct struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// OutputCreateProduct defines properties of input create product
type OutputCreateProduct struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
}

// OutputDeleteProduct defines properties of output delete product
type OutputDeleteProduct struct {
	Deleted bool   `json:"deleted"`
	Message string `json:"message"`
}

// OutputUpdateProduct defines properties of output update product
type OutputUpdateProduct struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
	Updated bool   `json:"updated"`
}

// OutputProduct defines properties of output find product
type OutputProduct struct {
	Message string            `json:"message"`
	Data    []*models.Product `json:"data"`
}

// OutputProducts defines properties of output find product
type OutputProducts struct {
	Message string            `json:"message"`
	Totals  int64             `json:"totals"`
	Data    []*models.Product `json:"data"`
}
