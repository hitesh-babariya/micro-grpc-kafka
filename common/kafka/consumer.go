// common/kafka/consumer.go
package kafka

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

func NewReader(topic, groupID string) *kafka.Reader {

	// ===== works local only not on kubernetes ========

	// return kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers: []string{"localhost:9092"},
	// 	Topic:   "user.created",
	// 	GroupID: "orders-service",
	// })

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	return kafka.NewReader(kafka.ReaderConfig{

		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset, // auto.offset.reset=earliest
	})

}

func Consume(ctx context.Context, reader *kafka.Reader, handler func([]byte)) {
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Kafka consume error:", err)
			continue
		}
		handler(msg.Value)
	}
}
