package handlers

import (
	"net/http"

	"github.com/berkayhellagun/microservice/src/product-api/data"
)

// swagger:route POST /products products addProduct
// Add new Product
//
// responses:
//	200: productResponse
// 422: errorValidation
// 501: errorResponse

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
