// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/product.go

// Package mock is a generated GoMock package.
package mock

import (
	storage "github.com/autsakorn/go-shop/storage"
	types "github.com/autsakorn/go-shop/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProduct is a mock of Product interface
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// CalSkip mocks base method
func (m *MockProduct) CalSkip(arg0, arg1 int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalSkip", arg0, arg1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// CalSkip indicates an expected call of CalSkip
func (mr *MockProductMockRecorder) CalSkip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalSkip", reflect.TypeOf((*MockProduct)(nil).CalSkip), arg0, arg1)
}

// CreateProduct mocks base method
func (m *MockProduct) CreateProduct(arg0 types.InputCreateProduct, arg1 storage.Product) (types.OutputCreateProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(types.OutputCreateProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct
func (mr *MockProductMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProduct)(nil).CreateProduct), arg0, arg1)
}

// DeleteProduct mocks base method
func (m *MockProduct) DeleteProduct(arg0 types.InputDeleteProduct, arg1 storage.Product) (types.OutputDeleteProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(types.OutputDeleteProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct
func (mr *MockProductMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProduct)(nil).DeleteProduct), arg0, arg1)
}

// FindProductByID mocks base method
func (m *MockProduct) FindProductByID(arg0 string, arg1 storage.Product) (types.OutputProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductByID", arg0, arg1)
	ret0, _ := ret[0].(types.OutputProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductByID indicates an expected call of FindProductByID
func (mr *MockProductMockRecorder) FindProductByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductByID", reflect.TypeOf((*MockProduct)(nil).FindProductByID), arg0, arg1)
}

// FindProducts mocks base method
func (m *MockProduct) FindProducts(arg0, arg1 int64, arg2 storage.Product) (types.OutputProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProducts", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.OutputProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProducts indicates an expected call of FindProducts
func (mr *MockProductMockRecorder) FindProducts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProducts", reflect.TypeOf((*MockProduct)(nil).FindProducts), arg0, arg1, arg2)
}

// UpdateProduct mocks base method
func (m *MockProduct) UpdateProduct(arg0 types.InputProduct, arg1 storage.Product) (types.OutputUpdateProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(types.OutputUpdateProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct
func (mr *MockProductMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProduct)(nil).UpdateProduct), arg0, arg1)
}

// UpsertProduct mocks base method
func (m *MockProduct) UpsertProduct(arg0 types.InputProduct, arg1 storage.Product) (types.OutputUpdateProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertProduct", arg0, arg1)
	ret0, _ := ret[0].(types.OutputUpdateProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertProduct indicates an expected call of UpsertProduct
func (mr *MockProductMockRecorder) UpsertProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertProduct", reflect.TypeOf((*MockProduct)(nil).UpsertProduct), arg0, arg1)
}