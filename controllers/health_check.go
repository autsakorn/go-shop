package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/autsakorn/go-shop/types"
	"github.com/julienschmidt/httprouter"
)

// HealthCheck route for check API
func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var results = types.OutputHealthCheck{
		Message: "API OK",
	}

	json.NewEncoder(w).Encode(results)
}
