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

import "github.com/berkayhellagun/microservice/src/product-api/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products return in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
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
