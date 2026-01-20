// orders/main.go
package main

import (
	"context"
	"kafka"
	"orders"
)

func main() {
	reader := kafka.NewReader()

	kafka.Consume(context.Background(), reader, orders.HandleUserCreated)
}
