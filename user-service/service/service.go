// user-service/service/service.go
package service

import (
	"context"
	"encoding/json"

	userspb "github.com/hitesh-babariya/micro-grpc-kafka/proto/user/v1"

	"github.com/hitesh-babariya/micro-grpc-kafka/common/kafka"
	kafkago "github.com/segmentio/kafka-go"

	"github.com/google/uuid"
)

type Service struct {
	userspb.UnimplementedUsersServiceServer
	writer *kafkago.Writer
}

func NewService(writer *kafkago.Writer) *Service {
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
	_ = kafka.Produce(ctx, s.writer, nil, event)

	return &user, nil
}
