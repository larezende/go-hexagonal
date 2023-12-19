package cli

import (
	"fmt"
	"go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "list":
		products, err := service.List()
		if err != nil {
			return result, err
		}
		for _, product := range products {
			result += fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n\n",
				product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
		}
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s %s has been enabled.", res.GetID(), res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s %s has been disabled.", res.GetID(), res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, nil
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())

	}

	return result, nil

}
