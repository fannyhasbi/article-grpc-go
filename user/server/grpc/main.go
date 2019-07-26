package main

import (
	"context"
	"github.com/fannyhasbi/article-grpc-go/user/services"
	"log"
	"net"

	pb "github.com/fannyhasbi/article-grpc-go/user/proto"
	"google.golang.org/grpc"
)

// UserServer implement UserServiceServer interface
type UserServer struct {
	Service *services.UserService
}

// FindUserByID find a user by ID
func (us UserServer) FindUserByID(ctx context.Context, ur *pb.UserRequest) (*pb.UserResponse, error) {
	result := us.Service.Query.FindUserByID(ur.Id)
	if result.Error != nil {
		return &pb.UserResponse{}, result.Error
	}
	if result.Result == nil {
		return &pb.UserResponse{}, result.Error
	}

	user := result.Result.(pb.UserResponse)

	return &user, nil
}

func main() {
	port := ":9000"
	srv := grpc.NewServer()
	var userServer UserServer

	userServer.Service = services.NewUserService()

	pb.RegisterUserServiceServer(srv, userServer)

	log.Println("Starting user RPC at", port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could not listen to %s: %v", port, err)
	}

	log.Fatal("User server is listening", srv.Serve(listen))

}
