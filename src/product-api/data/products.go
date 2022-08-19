package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// Product defines the structure for an api Product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name"`

	// the description for this product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	//configuration the validation func
	// first section is the field name and second section is validator function
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-abas-dsafa
	rege := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	match := rege.FindAllString(fl.Field().String(), -1)

	// if we dont have any match validation is gonna be fail
	if len(match) == 1 {
		return true
	}
	return false
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getId()
	productList = append(productList, p)
}

var ErrProductNotFound = fmt.Errorf("Prodcut not found.")

func FindProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

func GetProductById(id int) (*Product, error) {
	index := findIndexByProductID(id)
	if index == -1 {
		return nil, ErrProductNotFound
	}
	return productList[index], nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func getId() int {
	list := productList[len(productList)-1]
	return list.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       15.25,
		SKU:         "abc2121",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Without milk",
		Price:       10.25,
		SKU:         "zxc123",
	},
}
