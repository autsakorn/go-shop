package services

import (
	"errors"
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
		{
			name:       "Data create error case",
			data:       types.InputCreateProduct{Name: "Apple", Price: 1},
			createMock: func(product types.InputCreateProduct) error { return errors.New("Create Error") },
			want:       errors.New("Create Error"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		createMock = test.createMock
		_, err := CreateProduct(test.data, productStorage)

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
		{
			name:       "Data delete func error case",
			data:       types.InputDeleteProduct{ID: "1"},
			deleteMock: func(product types.InputDeleteProduct) error { return errors.New("Delete Error") },
			want:       errors.New("Delete Error"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		deleteMock = test.deleteMock
		_, err := DeleteProduct(test.data, productStorage)

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
		{
			name: "Database FindByID func error case",
			data: "1",
			findByIDMock: func(id string) (result []*models.Product, error error) {
				return result, errors.New("FindByID Error")
			},
			want: errors.New("FindByID Error"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		findByIDMock = test.findByIDMock
		_, err := FindProductByID(test.data, productStorage)

		if !reflect.DeepEqual(err, test.want) {
			t.Errorf("FindProductByID: %v fail  = %v, want %v", test.name, err, test.want)
		}
	}
}

func TestFindProducts(t *testing.T) {
	type inputPage struct {
		page  int64
		limit int64
	}
	tests := []struct {
		name      string
		data      inputPage
		findMock  func(skip int64, limit int64) ([]*models.Product, error)
		countMock func(product types.InputProduct) (int64, error)
		want      interface{}
	}{
		{
			name: "Happy case use default page and limit",
			data: inputPage{},
			findMock: func(skip int64, limit int64) ([]*models.Product, error) {
				var result []*models.Product
				return result, nil
			},
			countMock: func(product types.InputProduct) (int64, error) { return 1, nil },
			want:      nil,
		},
		{
			name: "Happy case",
			data: inputPage{page: 1, limit: 10},
			findMock: func(skip int64, limit int64) ([]*models.Product, error) {
				var result []*models.Product
				return result, nil
			},
			countMock: func(product types.InputProduct) (int64, error) { return 1, nil },
			want:      nil,
		},
		{
			name: "Request invalid page",
			data: inputPage{page: 2, limit: 10}, // Data criteria
			findMock: func(skip int64, limit int64) ([]*models.Product, error) {
				var result []*models.Product
				return result, nil
			},
			countMock: func(product types.InputProduct) (int64, error) { return 10, nil }, // Data criteria
			want:      errors.New("Invalid Page"),
		},
		{
			name: "Database error case",
			data: inputPage{page: 1, limit: 10},
			findMock: func(skip int64, limit int64) ([]*models.Product, error) {
				var result []*models.Product
				return result, errors.New("Database Error") // Return error
			},
			countMock: func(product types.InputProduct) (int64, error) { return 10, nil },
			want:      errors.New("Database Error"), // Expect error
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		findMock = test.findMock
		countMock = test.countMock
		_, err := FindProducts(test.data.page, test.data.limit, productStorage)
		if !reflect.DeepEqual(err, test.want) {
			t.Errorf("FindProducts: %v fail  = %v, want %v", test.name, err, test.want)
		}
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		name       string
		data       types.InputProduct
		updateMock func(product types.InputProduct) error
		want       interface{}
	}{
		{
			name:       "Happy case",
			data:       types.InputProduct{ID: "1"},
			updateMock: func(product types.InputProduct) error { return nil },
			want:       nil,
		},
		{
			name:       "Data update error case",
			data:       types.InputProduct{ID: "1"},
			updateMock: func(product types.InputProduct) error { return errors.New("Update Error") },
			want:       errors.New("Update Error"),
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		updateMock = test.updateMock
		_, err := UpdateProduct(test.data, productStorage)

		if !reflect.DeepEqual(err, test.want) {
			t.Errorf("UpdateProduct: %v fail = %v, want %v", test.name, err, test.want)
		}
	}
}

func TestUpsertProduct(t *testing.T) {
	type want struct {
		Err     interface{}
		Created bool
		Updated bool
	}
	tests := []struct {
		name       string
		data       types.InputProduct
		createMock func(product types.InputCreateProduct) error
		countMock  func(product types.InputProduct) (int64, error)
		updateMock func(product types.InputProduct) error
		want       want
	}{
		{
			name:       "Update case",
			data:       types.InputProduct{ID: "1"},
			createMock: func(product types.InputCreateProduct) error { return nil },
			countMock:  func(product types.InputProduct) (int64, error) { return 1, nil },
			updateMock: func(product types.InputProduct) error { return nil },
			want: want{
				Err:     nil,
				Created: false,
				Updated: true,
			},
		},
		{
			name:       "Create case",
			data:       types.InputProduct{ID: "1"},
			createMock: func(product types.InputCreateProduct) error { return nil },
			countMock:  func(product types.InputProduct) (int64, error) { return 0, nil },
			updateMock: func(product types.InputProduct) error { return nil },
			want: want{
				Err:     nil,
				Created: true,
				Updated: false,
			},
		},
		{
			name:       "Database Count func err case",
			data:       types.InputProduct{ID: "1"},
			createMock: func(product types.InputCreateProduct) error { return nil },
			countMock:  func(product types.InputProduct) (int64, error) { return 1, errors.New("Count Err") },
			updateMock: func(product types.InputProduct) error { return nil },
			want: want{
				Err:     errors.New("Count Err"),
				Created: false,
				Updated: false,
			},
		},
		{
			name:       "Database Update func err case",
			data:       types.InputProduct{ID: "1"},
			createMock: func(product types.InputCreateProduct) error { return nil },
			countMock:  func(product types.InputProduct) (int64, error) { return 1, nil },
			updateMock: func(product types.InputProduct) error { return errors.New("Update Err") },
			want: want{
				Err:     errors.New("Update Err"),
				Created: false,
				Updated: false,
			},
		},
		{
			name:       "Database Create func err case",
			data:       types.InputProduct{ID: "1"},
			createMock: func(product types.InputCreateProduct) error { return errors.New("Create Err") },
			countMock:  func(product types.InputProduct) (int64, error) { return 0, nil },
			updateMock: func(product types.InputProduct) error { return nil },
			want: want{
				Err:     errors.New("Create Err"),
				Created: false,
				Updated: false,
			},
		},
	}
	for _, test := range tests {
		productStorage := productStorageMock{}
		createMock = test.createMock
		countMock = test.countMock
		updateMock = test.updateMock
		result, err := UpsertProduct(test.data, productStorage)

		if !reflect.DeepEqual(err, test.want.Err) {
			t.Errorf("UpsertProduct: %v fail = %v, want err %v", test.name, err, test.want.Err)
		}
		if result.Created != test.want.Created {
			t.Errorf("UpsertProduct: %v create result = %v, want Created %v", test.name, result.Created, test.want.Created)
		}
		if result.Updated != test.want.Updated {
			t.Errorf("UpsertProduct: %v update result = %v, want Updated %v", test.name, result.Updated, test.want.Updated)
		}
	}
}
