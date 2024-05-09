package products

import (
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/repositories"
	"net/http"
)

func PaginateList(page int, itemsLimit int) (helpers.PaginateListResponse, error) {
	if itemsLimit == 0 {
		return helpers.PaginateListResponse{}, &helpers.GenericError{Msg: "itemsLimit cannot be equal to 0 ", Code: http.StatusUnprocessableEntity}
	}

	products, totalItems, err := repositories.ListPaginateProducts(page, itemsLimit)
	if err != nil {
		return helpers.PaginateListResponse{}, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	if products == nil {
		products = []models.Product{}
	}

	totalPages := (totalItems + int64(itemsLimit) - 1) / int64(itemsLimit)
	hasNextPage := int64(page) < totalPages

	response := helpers.PaginateListResponse{
		Data:        products,
		HasNextPage: hasNextPage,
		TotalPages:  totalPages,
	}

	return response, nil
}
