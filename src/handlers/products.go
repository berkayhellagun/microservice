package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect the id of product in url
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(group) != 1 {
			http.Error(rw, "Invalid Url", http.StatusBadRequest)
			return
		}
		if len(group[0]) != 2 {
			http.Error(rw, "Invalid Url", http.StatusBadRequest)
			return
		}
		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		}
		p.l.Println("Got Id", id)
		p.updateProduct(id, rw, r)
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	}
	// v is interface
	p.l.Printf("Prod %#v", prod)

	// go fake db
	errorProduct := data.UpdateProduct(id, prod)
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

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	}
	// v is interface
	p.l.Printf("Prod %#v", prod)

	// save the fake product list
	data.AddProduct(prod)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
