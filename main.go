package main

import (
	"context"
	"fmt"
	"github.com/jhuggart/repro-reflect-failure/gen/simple"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	port := 8080
	address := fmt.Sprintf(":%d", port)

	// Create gRPC server, register servers
	grpcServer := grpc.NewServer()

	simple.RegisterSimpleServiceServer(grpcServer, &simpleService{})

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("error creating listener", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve")
	}
}

type simpleService struct {
	simple.UnimplementedSimpleServiceServer
}

func (s *simpleService) GetSomething(_ context.Context, _ *simple.SomeQuery) (*simple.SomeResponse, error) {
	return &simple.SomeResponse{}, nil
}
