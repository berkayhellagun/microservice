package handlers

import (
	"net/http"

	"github.com/berkayhellagun/microservice/src/product-api/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// prod := &data.Product{}
	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	// }
	// v is interface
	p.l.Printf("Prod %#v", prod)

	// save the fake product list
	data.AddProduct(&prod)
}
