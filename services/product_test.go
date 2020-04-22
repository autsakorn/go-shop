package services

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/autsakorn/go-shop/models"
	"github.com/autsakorn/go-shop/types"
)

// Handle Mock Interface
var createMock func(product types.InputCreateProduct) error
var countMock func(product types.InputProduct) (int64, error)
var deleteMock func(product types.InputDeleteProduct) error
var findByIDMock func(id string) ([]*models.Product, error)
var findMock func(skip int64, limit int64) ([]*models.Product, error)
var updateMock func(product types.InputProduct) error
var upsertMock func(product types.InputProduct) error

type productStorageMock struct{}

func (s productStorageMock) Create(product types.InputCreateProduct) error {
	return createMock(product)
}
func (s productStorageMock) Count(product types.InputProduct) (int64, error) {
	return countMock(product)
}
func (s productStorageMock) Delete(product types.InputDeleteProduct) error {
	return deleteMock(product)
}
func (s productStorageMock) FindByID(id string) ([]*models.Product, error) {
	return findByIDMock(id)
}
func (s productStorageMock) Find(skip int64, limit int64) ([]*models.Product, error) {
	return findMock(skip, limit)
}
func (s productStorageMock) Update(product types.InputProduct) error {
	return updateMock(product)
}

// Start Test

func TestCalSkip(t *testing.T) {
	type input struct {
		page  int64
		limit int64
	}
	tests := []struct {
		name string
		data input
		want int64
	}{
		{name: "Page = 0", data: input{page: 0, limit: 5}, want: 0},
		{name: "Page = 1", data: input{page: 1, limit: 5}, want: 0},
		{name: "Page = 2", data: input{page: 2, limit: 5}, want: 5},
		{name: "Page = 5", data: input{page: 5, limit: 5}, want: 20},
		{name: "Page = 101", data: input{page: 101, limit: 5}, want: 500},
	}
	for _, test := range tests {
		skip := CalSkip(test.data.page, test.data.limit)
		log.Print(skip)
		if skip != test.want {
			t.Errorf("CalSkip: %v fail = %v, want %v", test.name, skip, test.want)
		}
	}
}
func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name       string
		data       types.InputCreateProduct
		createMock func(product types.InputCreateProduct) error
		want       interface{}
	}{
		{
			name:       "Happy case",
			data:       types.InputCreateProduct{Name: "Apple", Price: 100},
			createMock: func(product types.InputCreateProduct) error { return nil },
			want:       nil,
		},
		{
			name:       "Price required",
			data:       types.InputCreateProduct{Name: "Apple", Price: 0},
			createMock: func(product types.InputCreateProduct) error { return nil },
			want:       errors.New("Price is required"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		createMock = test.createMock
		_, err := CreateProduct(test.data, productStorage)
		log.Print(err)
		if !reflect.DeepEqual(err, test.want) {
			t.Errorf("CreateProduct %v fail = %v, want %v", test.name, err, test.want)
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		name       string
		data       types.InputDeleteProduct
		deleteMock func(product types.InputDeleteProduct) error
		want       interface{}
	}{
		{
			name:       "Happy case",
			data:       types.InputDeleteProduct{ID: "1"},
			deleteMock: func(product types.InputDeleteProduct) error { return nil },
			want:       nil,
		},
		{
			name:       "ID required",
			data:       types.InputDeleteProduct{},
			deleteMock: func(product types.InputDeleteProduct) error { return nil },
			want:       errors.New("ID is required"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		deleteMock = test.deleteMock
		_, err := DeleteProduct(test.data, productStorage)
		log.Print(err)
		if !reflect.DeepEqual(err, test.want) {
			t.Errorf("DeleteProduct: %v fail = %v, want %v", test.name, err, test.want)
		}
	}
}

func TestFindProductByID(t *testing.T) {
	tests := []struct {
		name         string
		data         string
		findByIDMock func(id string) ([]*models.Product, error)
		want         interface{}
	}{
		{
			name: "Happy case",
			data: "1",
			findByIDMock: func(id string) ([]*models.Product, error) {
				var result []*models.Product
				return result, nil
			},
			want: nil,
		},
		{
			name: "ID required case",
			data: "",
			findByIDMock: func(id string) ([]*models.Product, error) {
				var result []*models.Product
				return result, nil
			},
			want: errors.New("ID is required"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		findByIDMock = test.findByIDMock
		_, err := FindProductByID(test.data, productStorage)
		log.Print(err)
		if !reflect.DeepEqual(err, test.want) {
			t.Errorf("FindProductByID: %v fail  = %v, want %v", test.name, err, test.want)
		}
	}
}
