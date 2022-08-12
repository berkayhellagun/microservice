package handlers

import (
	"net/http"
	"strconv"

	"github.com/berkayhellagun/microservice/src/product-api/data"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// gorilla provide this struct
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // convert to int from url id(string)
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product) // we must cast the data

	// prod := &data.Product{}
	// err = prod.FromJSON(r.Body)

	// if err != nil {
	// 	http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	// }
	// // v is interface
	// p.l.Printf("Prod %#v", prod)

	// go fake db
	errorProduct := data.UpdateProduct(id, &prod)
	// check error status
	if errorProduct == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}
	if errorProduct != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
}
