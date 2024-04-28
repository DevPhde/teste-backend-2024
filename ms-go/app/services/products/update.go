package products

import (
	"ms-go/app/helpers"
	"ms-go/app/kafka/producers"
	"ms-go/app/models"
	"ms-go/app/repositories"
	"net/http"
	"time"
)

func Update(data models.Product, sendMessage bool) (*models.Product, error) {

	if data.ID == 0 {
		return nil, &helpers.GenericError{Msg: "Missing parameters", Code: http.StatusUnprocessableEntity}
	}

	product, err := repositories.GetProductById(data.ID)
	if err != nil {
		if product != nil {
			return nil, &helpers.GenericError{Msg: "Product Not Found", Code: http.StatusNotFound}
		}
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}
	var originalProduct = *product

	setUpdate(&data, product)

	if err := repositories.UpdateProduct(data); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	if sendMessage {
		err := producers.ProductProducer(data, "update")
		if err != nil {
			repositories.RollBackProduct(originalProduct)
			return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
		}
	}

	return &data, nil
}

func setUpdate(new, old *models.Product) {
	if new.ID == 0 {
		new.ID = old.ID
	}

	if new.Name == "" {
		new.Name = old.Name
	}

	if new.Brand == "" {
		new.Brand = old.Brand
	}

	if new.Price == 0 {
		new.Price = old.Price
	}

	if new.Description == "" {
		new.Description = old.Description
	}

	if new.Stock == 0 {
		new.Stock = old.Stock
	}

	new.CreatedAt = old.CreatedAt

	new.UpdatedAt = time.Now()
}
