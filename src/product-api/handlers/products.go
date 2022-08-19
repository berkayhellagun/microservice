package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

type KeyProduct struct{}
