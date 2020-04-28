package controllers

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHealthCheck_GetHealthCheck(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name string
		hc   HealthCheck
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := HealthCheck{}
			hc.GetHealthCheck(tt.args.w, tt.args.r, tt.args.in2)
		})
	}
}
