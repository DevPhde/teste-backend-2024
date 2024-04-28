package products

import (
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/repositories"
	"net/http"
)

func PaginateList(page int, itemsLimit int) (helpers.PaginateListResponse, error) {
	var response helpers.PaginateListResponse
	var products []models.Product

	if itemsLimit == 0 {
		return response, &helpers.GenericError{Msg: "itemsLimit cannot be equal to 0 ", Code: http.StatusUnprocessableEntity}
	}

	products, totalItems, err := repositories.ListPaginateProducts(page, itemsLimit)
	if err != nil {
		return response, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	totalPages := (totalItems + int64(itemsLimit) - 1) / int64(itemsLimit)

	hasNextPage := int64(page) < totalPages

	response.Data = products
	response.HasNextPage = hasNextPage
	response.TotalPages = totalPages

	return response, nil
}
