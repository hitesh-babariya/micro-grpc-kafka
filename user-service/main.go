// user-service/main.go
package main

import (
	"net"

	"kafka"
	"userspb"

	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	writer := kafka.NewWriter()
	service := NewService(writer)

	grpcServer := grpc.NewServer()
	userspb.RegisterUsersServiceServer(grpcServer, service)

	grpcServer.Serve(lis)
}
