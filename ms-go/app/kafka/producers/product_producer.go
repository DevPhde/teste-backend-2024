package producers

import (
	"context"
	"encoding/json"
	"errors"
	"ms-go/app/kafka"
	"ms-go/app/models"

	libKafka "github.com/segmentio/kafka-go"
)

func ProductProducer(data models.Product, productionType string) error {

	publisher, _ := kafka.KafkaConfiguration()

	jsonMessage, err := json.Marshal(data)
	if err != nil {
		return errors.New("internal error when sending message to broker")
	}
	err = publisher.WriteMessages(context.Background(), libKafka.Message{
		Key:   []byte(productionType),
		Value: jsonMessage,
	})

	if err != nil {
		return errors.New("internal error when sending message to broker")
	}

	publisher.Close()
	return nil
}
