package main

import (
	"fmt"
	"net"

	"github.com/NeverlandMJ/ToDo/todo-service/config"
	grpc_server "github.com/NeverlandMJ/ToDo/todo-service/grpc"
	"github.com/NeverlandMJ/ToDo/todo-service/server"
	"github.com/NeverlandMJ/ToDo/todo-service/service"
	"github.com/NeverlandMJ/ToDo/todo-service/todopb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	server, err := server.NewServer(cfg, "file://migrations")
	if err != nil {
		panic(err)
	}

	service := service.NewService(server)

	RunGRPCServer(*service)

}

// RunGRPCServer starts grpc server on localhost:9001
func RunGRPCServer(svc service.Service) {
	l, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}
	fmt.Println("server started at localhost:9001")
	s := grpc.NewServer()
	todopb.RegisterTodoServiceServer(s, grpc_server.NewgRPCServer(svc))

	if err = s.Serve(l); err != nil {
		panic(err)
	}
}
