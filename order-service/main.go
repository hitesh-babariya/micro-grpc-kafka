// orders/main.go
package main

import (
	"context"

	"github.com/hitesh-babariya/micro-grpc-kafka/common/kafka"
	"github.com/hitesh-babariya/micro-grpc-kafka/order-service/orders"
)

func main() {
	//reader := kafka.NewReader()
	reader := kafka.NewReader(
		"user.created",
		"order-service-group",
	)

	kafka.Consume(context.Background(), reader, orders.HandleUserCreated)
}
