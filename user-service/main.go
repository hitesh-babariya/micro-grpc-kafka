// user-service/main.go
package main

import (
	"net"

	"github.com/hitesh-babariya/micro-grpc-kafka/common/kafka"
	"github.com/hitesh-babariya/micro-grpc-kafka/user-service/service"

	userspb "github.com/hitesh-babariya/micro-grpc-kafka/proto/user/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	writer := kafka.NewWriter("user.created")
	service := service.NewService(writer)

	grpcServer := grpc.NewServer()
	userspb.RegisterUsersServiceServer(grpcServer, service)

	grpcServer.Serve(lis)
}
