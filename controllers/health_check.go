package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/autsakorn/go-shop/types"
	"github.com/julienschmidt/httprouter"
)

// HealthCheck Defines properties
type HealthCheck struct{}

// GetHealthCheck route for check API
func (hc HealthCheck) GetHealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var results = types.OutputHealthCheck{
		Message: "API OK",
	}

	json.NewEncoder(w).Encode(results)
}
