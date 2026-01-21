// common/kafka/producer.go
package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// NewWriter creates a Kafka writer (producer) for a given topic
func NewWriter(topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

// Produce sends a message to Kafka
func Produce(ctx context.Context, writer *kafka.Writer, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	err := writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Println("Failed to produce message:", err)
		return err
	}
	return nil
}
