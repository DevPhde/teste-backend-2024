package kafka

import (
	"github.com/segmentio/kafka-go"
)

func KafkaConfiguration() (*kafka.Writer, *kafka.Reader) {

	publisher := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:29092"},
		Topic:   "go-to-rails",
	})

	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:29092"},
		Topic:   "rails-to-go",
	})

	return publisher, consumer
}
