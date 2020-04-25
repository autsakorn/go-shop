package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/autsakorn/go-shop/models"
	"github.com/autsakorn/go-shop/storage/mock"
	"github.com/autsakorn/go-shop/types"
	"github.com/golang/mock/gomock"
)

func TestCalSkip(t *testing.T) {
	type args struct {
		page  int64
		limit int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "Page = 0", args: args{page: 0, limit: 5}, want: 0},
		{name: "Page = 1", args: args{page: 1, limit: 5}, want: 0},
		{name: "Page = 2", args: args{page: 2, limit: 5}, want: 5},
		{name: "Page = 5", args: args{page: 5, limit: 5}, want: 20},
		{name: "Page = 101", args: args{page: 101, limit: 5}, want: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalSkip(tt.args.page, tt.args.limit); got != tt.want {
				t.Errorf("CalSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
	type args struct {
		product                 types.InputCreateProduct
		mockCreateProductReturn error
	}
	tests := []struct {
		name    string
		args    args
		want    types.OutputCreateProduct
		wantErr bool
	}{
		{
			name:    "Create success case",
			args:    args{product: types.InputCreateProduct{Name: "Apple", Price: 100}, mockCreateProductReturn: nil},
			want:    types.OutputCreateProduct{Created: true, Message: ""},
			wantErr: false,
		},
		{
			name:    "Price required",
			args:    args{product: types.InputCreateProduct{Name: "Apple", Price: 0}, mockCreateProductReturn: nil},
			want:    types.OutputCreateProduct{Created: false, Message: "Price is required"},
			wantErr: true,
		},
		{
			name:    "Data create error case",
			args:    args{product: types.InputCreateProduct{Name: "Apple", Price: 1}, mockCreateProductReturn: errors.New("Create Error")},
			want:    types.OutputCreateProduct{Created: false, Message: "Product Create Error"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProduct := mock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Create(tt.args.product).
				AnyTimes().
				Return(tt.args.mockCreateProductReturn)

			got, err := CreateProduct(tt.args.product, mockProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	type args struct {
		product                 types.InputDeleteProduct
		mockDeleteProductReturn error
	}
	tests := []struct {
		name    string
		args    args
		want    types.OutputDeleteProduct
		wantErr bool
	}{
		{
			name:    "Success delete case",
			args:    args{product: types.InputDeleteProduct{ID: "1"}, mockDeleteProductReturn: nil},
			want:    types.OutputDeleteProduct{Deleted: true, Message: ""},
			wantErr: false,
		},
		{
			name:    "ID required",
			args:    args{product: types.InputDeleteProduct{ID: ""}, mockDeleteProductReturn: nil},
			want:    types.OutputDeleteProduct{Deleted: false, Message: "ID is required"},
			wantErr: true,
		},
		{
			name:    "DB delete func error case",
			args:    args{product: types.InputDeleteProduct{ID: "1"}, mockDeleteProductReturn: errors.New("Delete Error")},
			want:    types.OutputDeleteProduct{Deleted: false, Message: "Product Delete Error"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProduct := mock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Delete(tt.args.product).
				AnyTimes().
				Return(tt.args.mockDeleteProductReturn)

			got, err := DeleteProduct(tt.args.product, mockProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindProductByID(t *testing.T) {
	type mockFindProductReturn struct {
		data []*models.Product
		err  error
	}
	type args struct {
		id                    string
		mockFindProductReturn mockFindProductReturn
	}
	tests := []struct {
		name    string
		args    args
		want    types.OutputProduct
		wantErr bool
	}{
		{
			name:    "Success find product by ID",
			args:    args{id: "1", mockFindProductReturn: mockFindProductReturn{data: []*models.Product{{Name: "Product"}}, err: nil}},
			want:    types.OutputProduct{Message: "", Data: []*models.Product{{Name: "Product"}}},
			wantErr: false,
		},
		{
			name:    "ID Required case",
			args:    args{id: "", mockFindProductReturn: mockFindProductReturn{data: []*models.Product{{Name: "Product"}}, err: nil}},
			want:    types.OutputProduct{Message: "ID is required"},
			wantErr: true,
		},
		{
			name:    "Database FindByID func error case",
			args:    args{id: "2", mockFindProductReturn: mockFindProductReturn{data: []*models.Product{}, err: errors.New("FindByID Error")}},
			want:    types.OutputProduct{Message: "Product FindByID Error"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProduct := mock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				FindByID(tt.args.id).
				AnyTimes().
				Return(tt.args.mockFindProductReturn.data, tt.args.mockFindProductReturn.err)

			got, err := FindProductByID(tt.args.id, mockProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindProductByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindProducts(t *testing.T) {
	type mockFindProductInput struct {
		skip  int64
		limit int64
	}
	type mockFindProductReturn struct {
		data []*models.Product
		err  error
	}
	type mockCountProductReturn struct {
		result int64
		err    error
	}
	type args struct {
		page                   int64
		limit                  int64
		mockFindProductInput   mockFindProductInput
		mockFindProductReturn  mockFindProductReturn
		mockCountProductInput  types.InputProduct
		mockCountProductReturn mockCountProductReturn
	}
	tests := []struct {
		name    string
		args    args
		want    types.OutputProducts
		wantErr bool
	}{
		{
			name: "Success find product",
			args: args{
				mockFindProductInput:   mockFindProductInput{0, 10},
				mockFindProductReturn:  mockFindProductReturn{data: []*models.Product{{Name: "Product"}}, err: nil},
				mockCountProductInput:  types.InputProduct{},
				mockCountProductReturn: mockCountProductReturn{result: 1, err: nil},
			},
			want:    types.OutputProducts{Message: "", Totals: 1, Data: []*models.Product{{Name: "Product"}}},
			wantErr: false,
		},
		{
			name: "Request invalid page",
			args: args{
				page: 2, limit: 10,
				mockFindProductInput:   mockFindProductInput{10, 10},
				mockFindProductReturn:  mockFindProductReturn{data: []*models.Product{{Name: "Product"}}, err: nil},
				mockCountProductInput:  types.InputProduct{},
				mockCountProductReturn: mockCountProductReturn{result: 1, err: nil},
			},
			want:    types.OutputProducts{Message: "Invalid Page", Totals: 1},
			wantErr: true,
		},
		{
			name: "Find product database error",
			args: args{
				page: 1, limit: 10,
				mockFindProductInput:   mockFindProductInput{0, 10},
				mockFindProductReturn:  mockFindProductReturn{data: []*models.Product{}, err: errors.New("Find Error")},
				mockCountProductInput:  types.InputProduct{},
				mockCountProductReturn: mockCountProductReturn{result: 1, err: nil},
			},
			want:    types.OutputProducts{Message: "Product Find Error", Totals: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProduct := mock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Count(tt.args.mockCountProductInput).
				AnyTimes().
				Return(tt.args.mockCountProductReturn.result, tt.args.mockCountProductReturn.err)

			mockProduct.EXPECT().
				Find(tt.args.mockFindProductInput.skip, tt.args.mockFindProductInput.limit).
				AnyTimes().
				Return(tt.args.mockFindProductReturn.data, tt.args.mockFindProductReturn.err)

			got, err := FindProducts(tt.args.page, tt.args.limit, mockProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	type args struct {
		product                 types.InputProduct
		mockUpdateProductReturn error
	}
	tests := []struct {
		name    string
		args    args
		want    types.OutputUpdateProduct
		wantErr bool
	}{
		{
			name: "Success update product",
			args: args{
				product:                 types.InputProduct{ID: "1"},
				mockUpdateProductReturn: nil,
			},
			want: types.OutputUpdateProduct{
				Updated: true,
				Created: false,
			},
			wantErr: false,
		},
		{
			name: "Update product database error",
			args: args{
				product:                 types.InputProduct{ID: "1"},
				mockUpdateProductReturn: errors.New("Update Error"),
			},
			want: types.OutputUpdateProduct{
				Updated: false,
				Created: false,
				Message: "Product Update Error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProduct := mock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Update(tt.args.product).
				AnyTimes().
				Return(tt.args.mockUpdateProductReturn)

			got, err := UpdateProduct(tt.args.product, mockProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpsertProduct(t *testing.T) {
	type mockOutputCountProduct struct {
		numberProduct int64
		err           error
	}
	type args struct {
		product                 types.InputProduct
		mockOutputCountProduct  mockOutputCountProduct
		mockOutputCreateProduct error
		mockOutputUpdateProduct error
		mockInputCreateProduct  types.InputCreateProduct
	}
	tests := []struct {
		name    string
		args    args
		want    types.OutputUpdateProduct
		wantErr bool
	}{
		{
			name: "Success update product",
			args: args{
				product:                 types.InputProduct{ID: "1", Name: "Product"},
				mockOutputCountProduct:  mockOutputCountProduct{1, nil},
				mockOutputCreateProduct: nil,
				mockOutputUpdateProduct: nil,
			},
			want: types.OutputUpdateProduct{
				Created: false,
				Updated: true,
			},
			wantErr: false,
		},
		{
			name: "Success create product",
			args: args{
				product:                 types.InputProduct{ID: "2", Name: "Product New"},
				mockInputCreateProduct:  types.InputCreateProduct{Name: "Product New"},
				mockOutputCountProduct:  mockOutputCountProduct{0, nil},
				mockOutputCreateProduct: nil,
				mockOutputUpdateProduct: nil,
			},
			want: types.OutputUpdateProduct{
				Created: true,
				Updated: false,
			},
			wantErr: false,
		},
		{
			name: "Count database error",
			args: args{
				product:                 types.InputProduct{ID: "2", Name: "Product New"},
				mockInputCreateProduct:  types.InputCreateProduct{Name: "Product New"},
				mockOutputCountProduct:  mockOutputCountProduct{0, errors.New("Count Error")},
				mockOutputCreateProduct: nil,
				mockOutputUpdateProduct: nil,
			},
			want: types.OutputUpdateProduct{
				Created: false,
				Updated: false,
				Message: "Product Count Error",
			},
			wantErr: true,
		},
		{
			name: "Update database error",
			args: args{
				product:                 types.InputProduct{ID: "1", Name: "Product"},
				mockOutputCountProduct:  mockOutputCountProduct{1, nil},
				mockOutputCreateProduct: nil,
				mockOutputUpdateProduct: errors.New("Update Error"),
			},
			want: types.OutputUpdateProduct{
				Created: false,
				Updated: false,
				Message: "Product Update Error",
			},
			wantErr: true,
		},
		{
			name: "Create database error",
			args: args{
				product:                 types.InputProduct{ID: "2", Name: "Product New"},
				mockInputCreateProduct:  types.InputCreateProduct{Name: "Product New"},
				mockOutputCountProduct:  mockOutputCountProduct{0, nil},
				mockOutputCreateProduct: errors.New("Create Error"),
				mockOutputUpdateProduct: nil,
			},
			want: types.OutputUpdateProduct{
				Created: false,
				Updated: false,
				Message: "Product Create Error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockProduct := mock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Count(tt.args.product).
				AnyTimes().
				Return(tt.args.mockOutputCountProduct.numberProduct, tt.args.mockOutputCountProduct.err)
			mockProduct.EXPECT().
				Create(tt.args.mockInputCreateProduct).
				AnyTimes().
				Return(tt.args.mockOutputCreateProduct)
			mockProduct.EXPECT().
				Update(tt.args.product).
				AnyTimes().
				Return(tt.args.mockOutputUpdateProduct)

			got, err := UpsertProduct(tt.args.product, mockProduct)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpsertProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpsertProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
