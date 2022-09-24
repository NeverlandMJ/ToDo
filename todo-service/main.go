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

	server, err := server.NewServer(cfg)
	if err != nil {
		panic(err)
	}

	service := service.NewService(server)

	RunGRPCServer(*service, cfg)

}

// RunGRPCServer starts grpc server on localhost:9001
func RunGRPCServer(svc service.Service, cfg config.Config) {
	l, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	todopb.RegisterTodoServiceServer(s, grpc_server.NewgRPCServer(svc))
	
	fmt.Println("Server started at ", l.Addr())
	
	if err = s.Serve(l); err != nil {
		panic(err)
	}
}
