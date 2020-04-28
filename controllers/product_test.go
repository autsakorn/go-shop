package controllers

import (
	"net/http"
	"testing"

	"github.com/autsakorn/go-shop/services"
	"github.com/autsakorn/go-shop/storage"
	"github.com/julienschmidt/httprouter"
)

func TestProduct_CreateProduct(t *testing.T) {
	type fields struct {
		ProductService services.Product
		Storage        storage.Storage
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ProductService: tt.fields.ProductService,
				Storage:        tt.fields.Storage,
			}
			p.CreateProduct(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}

func TestProduct_DeleteProduct(t *testing.T) {
	type fields struct {
		ProductService services.Product
		Storage        storage.Storage
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ProductService: tt.fields.ProductService,
				Storage:        tt.fields.Storage,
			}
			p.DeleteProduct(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}

func TestProduct_FindProducts(t *testing.T) {
	type fields struct {
		ProductService services.Product
		Storage        storage.Storage
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ProductService: tt.fields.ProductService,
				Storage:        tt.fields.Storage,
			}
			p.FindProducts(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}

func TestProduct_FindProductByID(t *testing.T) {
	type fields struct {
		ProductService services.Product
		Storage        storage.Storage
	}
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ProductService: tt.fields.ProductService,
				Storage:        tt.fields.Storage,
			}
			p.FindProductByID(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}

func TestProduct_UpdateProduct(t *testing.T) {
	type fields struct {
		ProductService services.Product
		Storage        storage.Storage
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ProductService: tt.fields.ProductService,
				Storage:        tt.fields.Storage,
			}
			p.UpdateProduct(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}

func TestProduct_UpsertProduct(t *testing.T) {
	type fields struct {
		ProductService services.Product
		Storage        storage.Storage
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ProductService: tt.fields.ProductService,
				Storage:        tt.fields.Storage,
			}
			p.UpsertProduct(tt.args.w, tt.args.req, tt.args.in2)
		})
	}
}
