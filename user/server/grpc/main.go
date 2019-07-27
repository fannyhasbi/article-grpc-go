package main

import (
	"log"
	"net"

	pb "github.com/fannyhasbi/article-grpc-go/user/proto"
	"github.com/fannyhasbi/article-grpc-go/user/server/grpc/handler"
	"github.com/fannyhasbi/article-grpc-go/user/services"
	"google.golang.org/grpc"
)

func main() {
	port := ":9000"
	srv := grpc.NewServer()
	var userServer handler.UserServer

	userServer.Service = services.NewUserService()

	pb.RegisterUserServiceServer(srv, userServer)

	log.Println("Starting user RPC at", port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could not listen to %s: %v", port, err)
	}

	log.Fatal("User server is listening", srv.Serve(listen))

}
