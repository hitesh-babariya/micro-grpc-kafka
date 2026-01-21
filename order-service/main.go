// orders/main.go
package main

import (
	"context"

	"github.com/hitesh-babariya/micro-grpc-kafka/common/kafka"
	"github.com/hitesh-babariya/micro-grpc-kafka/order-service/orders"
)

func main() {
	reader := kafka.NewReader()

	kafka.Consume(context.Background(), reader, orders.HandleUserCreated)
}
