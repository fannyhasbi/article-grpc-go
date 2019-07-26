package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fannyhasbi/article-grpc-go/user/proto"
	"github.com/fannyhasbi/article-grpc-go/user/storage"
	"google.golang.org/grpc"
)

// UserServer implement UserServiceServer interface
type UserServer struct{}

// FindUserByID find a user by ID
func (us UserServer) FindUserByID(ctx context.Context, ur *pb.UserRequest) (*pb.UserResponse, error) {
	var response *pb.UserResponse

	storage := storage.NewUserStorage()

	for _, v := range storage.UserMap {
		if v.ID == ur.Id {
			response = &pb.UserResponse{
				Id:    v.ID,
				Name:  v.Name,
				Email: v.Email,
			}

			return response, nil
		}
	}
	return &pb.UserResponse{}, nil
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
