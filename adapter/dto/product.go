package dto

import "github.com/nironwp/hexagonal-architecture/application"

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
	Quantity int     `json:"quantity"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	err := product.IsValid()

	if err != nil {
		return nil, err
	}

	return product, nil
}
