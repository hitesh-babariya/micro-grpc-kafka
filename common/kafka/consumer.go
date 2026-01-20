// common/kafka/consumer.go
package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user.created",
		GroupID: "orders-service",
	})
}

func Consume(ctx context.Context, reader *kafka.Reader, handler func([]byte)) {
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		handler(msg.Value)
	}
}
