package repositories

import (
	"context"
	"errors"
	"ms-go/app/models"
	"ms-go/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListPaginateProducts(page, itemsLimit int) ([]models.Product, int64, error) { // Get list of products with pagination
	var products []models.Product

	skip := (page - 1) * itemsLimit

	totalItems, err := db.Connection().CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return nil, 0, errors.New("internal Error. Try again later")
	}

	data, err := db.Connection().Find(context.TODO(), bson.D{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsLimit)))
	if err != nil {
		return nil, 0, errors.New("internal Error. Try again later")
	}

	for data.Next(context.TODO()) {
		var product models.Product
		err := data.Decode(&product)
		if err != nil {
			return nil, 0, errors.New("internal Error. Try again later")
		}
		products = append(products, product)
	}

	return products, totalItems, nil
}

func GetProductNextId(ctx context.Context) (int, error) { // Return new Product id of collection
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	var max models.Product
	if err := db.Connection().FindOne(ctx, bson.D{}, opts).Decode(&max); err != nil {
		if err == mongo.ErrNoDocuments {
			return 1, nil
		}
		return 0, errors.New("internal Error. Try again later")
	}
	return max.ID + 1, nil
}

func Create(product models.Product) (*models.Product, error) { // Create new Product
	_, err := db.Connection().InsertOne(context.TODO(), product)

	if err != nil {
		return nil, errors.New("internal Error. Try again later")
	}
	defer db.Disconnect()
	return &product, nil
}

func GetProductById(productId int) (*models.Product, error) { // get a product by id
	var product models.Product
	if err := db.Connection().FindOne(context.TODO(), bson.M{"id": productId}).Decode(&product); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &product, errors.New("product Not Found")
		}
		return nil, errors.New("internal Error. Try again later")
	}

	defer db.Disconnect()
	return &product, nil
}

func UpdateProduct(product models.Product) error {

	if err := db.Connection().FindOneAndUpdate(context.TODO(), bson.M{"id": product.ID}, bson.M{"$set": product}).Decode(&product); err != nil {
		return errors.New("internal Error. Try again later")
	}
	return nil
}

func SecurityRemovalProduct(productId int) error { // used when kafka broker failed, removing the product of database to previne inconsistencies
	filter := bson.D{{Key: "id", Value: productId}}

	_, err := db.Connection().DeleteOne(context.Background(), filter)
	if err != nil {
		return errors.New("internal Error. Try again later")
	}

	return nil
}

func RollBackProduct(product models.Product) error {
	if err := db.Connection().FindOneAndUpdate(context.TODO(), bson.M{"id": product.ID}, bson.M{"$set": product}).Decode(&product); err != nil {
		return errors.New("internal Error. Try again later")
	}
	return nil
}
