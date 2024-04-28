package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-go/app/kafka"
	"ms-go/app/models"
	"ms-go/app/services/products"
)

func ProductConsumer() {
	_, consumer := kafka.KafkaConfiguration()
	for {
		message, err := consumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error when consuming topic:", err)
			continue
		}

		key := string(message.Key)
		var product models.Product

		if err := json.Unmarshal(message.Value, &product); err != nil {
			fmt.Println("Error when trying to unmarshal message:", err)
			continue
		}
		fmt.Println("Received message value:", "key: ", key, "message: ", product)

		switch key {
		case "create":
			products.Create(product, false)
			fmt.Println("Processing 'create' message...")
		case "update":
			products.Update(product, false)
			fmt.Println("Processing 'update' message...")
		default:
			fmt.Println("Unexpected key in message:", key)
		}
	}
}
