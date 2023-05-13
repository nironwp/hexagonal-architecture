package cli

import (
	"fmt"

	"github.com/nironwp/hexagonal-architecture/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64,
	quantity int,
) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price, quantity)

		if err != nil {
			return result, err
		}
		product.GetID()
		result = fmt.Sprintf(
			"Product ID %s with the name %s has been created with price %f, quantity of %b and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetQuantity(), product.GetStatus(),
		)
	case "enable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product with ID %s, and name %s is now %s", res.GetID(), res.GetName(), res.GetStatus())

	case "disable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product with ID %s, and name %s is now %s", res.GetID(), res.GetName(), res.GetStatus())
	default:
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product ID: %s\n Name: %s\n Price: %f\n Quantity: %b\n Status: %s\n ", product.GetID(), product.GetName(), product.GetPrice(), product.GetQuantity(), product.GetStatus())
	}

	return result, nil
}
