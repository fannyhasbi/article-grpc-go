package main

import (
	"context"
	"log"

	"github.com/fannyhasbi/article-grpc-go/article/config"
	pb "github.com/fannyhasbi/article-grpc-go/article/proto"
)

func main() {
	userRequest := pb.UserRequest{
		Id: 123,
	}

	grpc := config.InitGRPCConnection()

	result, err := grpc.FindUserByID(context.Background(), &userRequest)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%#v\n", result)

}
