package handlers

import (
	"net/http"

	"github.com/berkayhellagun/microservice/src/product-api/data"
)

type productResponseWrapper struct {
	// all product in system
	Body []data.Product
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	listOfProduct := data.GetProducts()
	// encode as JSON

	// d, err := json.Marshal(listOfProduct)
	// we can use marshal but encoder is more powerfull than marshal so we use encoder
	err := listOfProduct.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encoder json", http.StatusInternalServerError)
	}
	// write the data
	//rw.Write(d)
}
