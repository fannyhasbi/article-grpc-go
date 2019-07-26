package config

import (
	"log"

	pb "github.com/fannyhasbi/article-grpc-go/article/proto"
	"google.golang.org/grpc"
)

// InitGRPCConnection init a user GRPC client
func InitGRPCConnection() pb.UserServiceClient {
	host := "localhost:9000"

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", host, err)
	}

	return pb.NewUserServiceClient(conn)
}
