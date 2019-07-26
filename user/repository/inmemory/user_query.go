package inmemory

import (
	pb "github.com/fannyhasbi/article-grpc-go/user/proto"
	"github.com/fannyhasbi/article-grpc-go/user/repository"
	"github.com/fannyhasbi/article-grpc-go/user/storage"
)

// UserQueryInMemory implement UserQuery interface
type UserQueryInMemory struct {
	Storage *storage.UserStorage
}

// NewUserQueryInMemory create user query instance
func NewUserQueryInMemory(storage *storage.UserStorage) repository.UserQuery {
	return &UserQueryInMemory{Storage: storage}
}

// FindUserByID find a user by id
func (uq *UserQueryInMemory) FindUserByID(uid int32) repository.QueryResult {
	var response repository.QueryResult

	for _, v := range uq.Storage.UserMap {
		if v.ID == uid {
			response.Result = pb.UserResponse{
				Id:    v.ID,
				Name:  v.Name,
				Email: v.Email,
			}

			return response
		}
	}

	return response
}
