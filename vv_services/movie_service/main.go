package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"movie_service/internal/app/di"
	"movie_service/internal/app/interface/rpc/v1/protocol"
	"net"
)

const (
	// todo set from env
	port = ":9090"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	service := di.InitializeMovieService()
	protocol.RegisterMovieServiceServer(server, &service)

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
