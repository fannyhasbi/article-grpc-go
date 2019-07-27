package handler

import (
	"context"

	pb "github.com/fannyhasbi/article-grpc-go/user/proto"
	"github.com/fannyhasbi/article-grpc-go/user/services"
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
