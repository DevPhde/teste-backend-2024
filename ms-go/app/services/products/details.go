package products

import (
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/repositories"
	"net/http"
)

func Datails(productId int) (*models.Product, error) {

	if productId == 0 {
		return nil, &helpers.GenericError{Msg: "Missing params", Code: http.StatusBadRequest}
	}

	product, err := repositories.GetProductById(productId)
	if err != nil {
		if product != nil {
			return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusNotFound}
		}
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	return product, nil
}
