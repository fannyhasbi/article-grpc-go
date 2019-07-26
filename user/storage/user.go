package storage

import (
	"github.com/fannyhasbi/article-grpc-go/user/domain"
)

type UserStorage struct {
	UserMap []domain.User
}

var userMap = []domain.User{
	domain.User{
		ID:    121,
		Name:  "User 1",
		Email: "user1@example.com",
	},
	domain.User{
		ID:    122,
		Name:  "User 2",
		Email: "user2@example.com",
	},
	domain.User{
		ID:    123,
		Name:  "User 3",
		Email: "user3@example.com",
	},
	domain.User{
		ID:    124,
		Name:  "User 4",
		Email: "user4@example.com",
	},
}

// NewUserStorage create new user storage instance
func NewUserStorage() *UserStorage {
	return &UserStorage{
		UserMap: userMap,
	}
}
