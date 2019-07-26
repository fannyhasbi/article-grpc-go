package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fannyhasbi/article-grpc-go/user/proto"
	"google.golang.org/grpc"
)

// UserServer implement UserServiceServer interface
type UserServer struct{}

// FindUserByID find a user by ID
func (us UserServer) FindUserByID(ctx context.Context, ur *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:    1234,
		Name:  "Fanny Hasbi",
		Email: "hasbi@example.com",
	}, nil
}

func main() {
	port := ":9000"
	srv := grpc.NewServer()
	var userServer UserServer

	pb.RegisterUserServiceServer(srv, userServer)

	log.Println("Starting user RPC at", port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could not listen to %s: %v", port, err)
	}

	log.Fatal("User server is listening", srv.Serve(listen))

}
