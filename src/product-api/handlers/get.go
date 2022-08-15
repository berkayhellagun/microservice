package handlers

import (
	"net/http"

	"github.com/berkayhellagun/microservice/src/product-api/data"
)

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	listOfProduct := data.GetProducts()
	// encode as JSON

	// d, err := json.Marshal(listOfProduct)
	// we can use marshal but encoder is more powerfull than marshal so we use encoder
	err := data.ToJSON(listOfProduct, rw)
	if err != nil {
		http.Error(rw, "Unable to encoder json", http.StatusInternalServerError)
	}
	// write the data
	//rw.Write(d)
}
