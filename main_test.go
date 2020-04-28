// main.go
package main

import (
	"testing"

	"github.com/autsakorn/go-shop/controllers"
)

func Test_routes(t *testing.T) {
	type args struct {
		c controllers.Controllers
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routes(tt.args.c)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
