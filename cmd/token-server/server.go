package main

import (
	"log"
	"net"
	"todo-list/internal"
	"todo-list/protogen/token"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()
	tokenServer := internal.NewTokenServer()
	token.RegisterTokenServer(grpcServer, &tokenServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server", err)
	}
}
