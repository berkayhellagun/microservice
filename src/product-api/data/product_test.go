package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "press",
		Price: 10,
		SKU:   "test-test-test",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
