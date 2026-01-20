// orders/consumer.go
package orders

import (
	"encoding/json"
	"log"
)

type UserCreatedEvent struct {
	Id    string
	Name  string
	Email string
}

func HandleUserCreated(data []byte) {
	var event UserCreatedEvent
	json.Unmarshal(data, &event)

	log.Println("User received in Orders:", event.Id)

	// Example:
	// - initialize user order history
	// - cache user data
}
