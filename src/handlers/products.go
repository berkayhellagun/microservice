package handlers

import (
	"log"
	"net/http"

	"github.com/berkayhellagun/microservice/src/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
