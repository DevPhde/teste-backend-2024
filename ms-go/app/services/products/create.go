package products

import (
	"context"
	"net/http"
	"time"

	"ms-go/app/helpers"
	"ms-go/app/kafka/producers"
	"ms-go/app/models"
	"ms-go/app/repositories"
)

func Create(data models.Product, sendMessage bool) (*models.Product, error) {

	if data.ID == 0 {
		id, err := repositories.GetProductNextId(context.TODO())
		if err != nil {
			return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
		}

		data.ID = id
	}

	if err := data.Validate(); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusUnprocessableEntity}
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = data.CreatedAt

	if _, err := repositories.Create(data); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	if sendMessage {
		productCopy := data
		err := producers.ProductProducer(productCopy, "create")
		if err != nil {
			repositories.SecurityRemovalProduct(data.ID)
			return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
		}
	}

	return &data, nil
}
