package controllers

import (
	"encoding/json"
	"net/http"
	"log"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

func GetProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := map[string]int{"example": 7}
	json.NewEncoder(w).Encode(result)
}

func GetProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	log.Print(id)
	result := map[string]int{"example": 7}
	json.NewEncoder(w).Encode(result)
}

func PostProduct(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	result := map[string]int{"example": 7}
	json.NewEncoder(w).Encode(result)
}

func PatchProduct(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	result := map[string]int{"example": 7}
	json.NewEncoder(w).Encode(result)
}