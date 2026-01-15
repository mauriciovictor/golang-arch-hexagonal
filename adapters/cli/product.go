package cli

import (
	"fmt"

	"github.com/mauriciovictor/curso-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s  with price %f and status %s  was created successfully", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s and price %f was enabled successfully", res.GetID(), res.GetName(), res.GetPrice())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s and price %f was disabled successfully", res.GetID(), res.GetName(), res.GetPrice())

	default:
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s and price %f was found successfully", product.GetID(), product.GetName(), product.GetPrice())
	}

	return result, nil
}
