package services

import (
	"github.com/fannyhasbi/article-grpc-go/user/repository"
	"github.com/fannyhasbi/article-grpc-go/user/repository/inmemory"
	"github.com/fannyhasbi/article-grpc-go/user/storage"
)

// UserService wrap user services
type UserService struct {
	Query repository.UserQuery
}

// NewUserService create new UserService instance
func NewUserService() *UserService {
	var userQuery repository.UserQuery

	db := storage.NewUserStorage()
	userQuery = inmemory.NewUserQueryInMemory(db)

	return &UserService{
		Query: userQuery,
	}
}
