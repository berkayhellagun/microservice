package handlers

import (
	"net/http"

	"github.com/berkayhellagun/microservice/src/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
//
// responses:
//	200: productsResponse

func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET List of Product")
	prod := data.GetProducts()
	// convert to json
	err := data.ToJSON(prod, rw)
	if err != nil {
		p.l.Fatal("Unable to decoder json")
		http.Error(rw, "Unable to decoder json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
//
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests

func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	prod, err := data.GetProductById(id)
	if err != nil {
		p.l.Fatal("Product does not exist!")
	}
	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Fatal("Unable to decoder json")
		http.Error(rw, "Unable to decoder json", http.StatusInternalServerError)
	}
}

// func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
// 	p.l.Println("Handle GET Products")
// 	listOfProduct := data.GetProducts()
// 	// encode as JSON

// 	// d, err := json.Marshal(listOfProduct)
// 	// we can use marshal but encoder is more powerfull than marshal so we use encoder
// 	err := data.ToJSON(listOfProduct, rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to encoder json", http.StatusInternalServerError)
// 	}
// 	// write the data
// 	//rw.Write(d)
// }
