package grpcClient

import (
	"token-management-service/protogen/token"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateGrpcClient() (token.TokenClient, *grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}
	// defer conn.Close()

	t := token.NewTokenClient(conn)
	return t, conn, nil
}
