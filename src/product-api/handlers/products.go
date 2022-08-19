// Package classification of Product API
//
// # Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Verion: 1.0.0
//
//	Consumes:
//	-application/json
//
//	Produces:
//	-application/json
//
//swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/berkayhellagun/microservice/src/product-api/data"
	"github.com/gorilla/mux"
)

// A list of products return in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

//		if r.Method == http.MethodPut {
//			// expect the id of product in url
//			regex := regexp.MustCompile(`/([0-9]+)`)
//			group := regex.FindAllStringSubmatch(r.URL.Path, -1)
//			if len(group) != 1 {
//				http.Error(rw, "Invalid Url", http.StatusBadRequest)
//				return
//			}
//			if len(group[0]) != 2 {
//				http.Error(rw, "Invalid Url", http.StatusBadRequest)
//				return
//			}
//			idString := group[0][1]
//			id, err := strconv.Atoi(idString)
//			if err != nil {
//				http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
//			}
//			p.l.Println("Got Id", id)
//			p.updateProduct(id, rw, r)
//		}
//		//catch all
//		rw.WriteHeader(http.StatusMethodNotAllowed)
//	}
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := data.FromJSON(prod, r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("[ERROR] validating product %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
