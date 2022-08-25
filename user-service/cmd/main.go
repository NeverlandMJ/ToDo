package main

import (
	"net"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	grpc_server "github.com/NeverlandMJ/ToDo/user-service/grpc"
	"github.com/NeverlandMJ/ToDo/user-service/server"
	"github.com/NeverlandMJ/ToDo/user-service/service"
	"github.com/NeverlandMJ/ToDo/user-service/v1/userpb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	server, err := server.NewServer(cfg)
	if err != nil {
		panic(err)
	}

	service := service.NewService(server)

	RunGRPCServer(*service)

}

func RunGRPCServer(svc service.Service) {
	l, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, grpc_server.NewgRPCServer(svc))

	if err = s.Serve(l); err != nil {
		panic(err)
	}
}
