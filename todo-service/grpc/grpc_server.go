package grpc

import (
	"github.com/NeverlandMJ/ToDo/todo-service/service"
	"github.com/NeverlandMJ/ToDo/todo-service/todopb"
)

type gRPCServer struct {
	todopb.UnimplementedTodoServiceServer
	svc service.Service
}

func NewgRPCServer(svc service.Service) *gRPCServer {
	return &gRPCServer{
		svc: svc,
	}
}

