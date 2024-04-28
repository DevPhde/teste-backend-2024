package helpers

import "ms-go/app/models"

type PaginateListResponse struct {
	Data        []models.Product
	HasNextPage bool
	TotalPages  int64
}
