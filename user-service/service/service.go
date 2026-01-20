// user-service/service/service.go
package service

import (
	"context"
	"encoding/json"

	"userspb"

	"github.com/hitesh-babariya/common/kafka"

	"github.com/google/uuid"
)

type Service struct {
	userspb.UnimplementedUsersServiceServer
	writer *kafka.Writer
}

func NewService(writer *kafka.Writer) *Service {
	return &Service{writer: writer}
}

func (s *Service) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.UserResponse, error) {
	id := uuid.New().String()

	user := userspb.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	// publish event
	event, _ := json.Marshal(user)
	_ = kafka.Publish(ctx, s.writer, event)

	return &user, nil
}
