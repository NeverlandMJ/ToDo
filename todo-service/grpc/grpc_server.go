package grpc

import (
	"context"
	"errors"
	"log"
	"time"

	customerr "github.com/NeverlandMJ/ToDo/todo-service/pkg/customERR"
	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/todo-service/service"
	"github.com/NeverlandMJ/ToDo/todo-service/todopb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
    layout = "2006-01-02"
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

func (g *gRPCServer) CreateTodo(ctx context.Context, req *todopb.RequestTodo) (*todopb.ResponseTodo, error) {
	id, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "UserID is not uuid")
	}

	ddl, err := time.Parse(layout, req.Deadline)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "Deadline is not parsable")
	}

	td, err := g.svc.CreateTodo(ctx, entity.Todo{
		UserID:   id,
		Body:     req.GetBody(),
		Deadline: ddl,
	})

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "couldn't create todo")
	}

	return &todopb.ResponseTodo{
		Id:     td.ID.String(),
		UserID: td.UserID.String(),
		Body: td.Body,
		CreatedAt: td.CreatedAt.String(),
		Deadline: td.Deadline.String(),
		IsDone: td.IsDone,
	}, nil
}

func (g *gRPCServer) GetTodoByID(ctx context.Context, req *todopb.RequestTodoID) (*todopb.ResponseTodo, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	td, err := g.svc.GetTodo(ctx, id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, customerr.ERR_TODO_NOT_EXIST) {
			return nil, status.Error(codes.NotFound, "todo with given id is not found")
		}else {
			return nil, status.Error(codes.Internal, "couldn't fetch todo data")
		}
	}

	return &todopb.ResponseTodo{
		Id:     td.ID.String(),
		UserID: td.UserID.String(),
		Body: td.Body,
		CreatedAt: td.CreatedAt.String(),
		Deadline: td.Deadline.String(),
		IsDone: td.IsDone,
	}, nil

}

func (g *gRPCServer) MarkAsDone(ctx context.Context, req *todopb.RequestTodoID) (*todopb.Empty, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	err = g.svc.MarkAsDone(ctx, id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, customerr.ERR_TODO_NOT_EXIST) {
			return nil, status.Error(codes.NotFound, "todo with given id is not found")
		}else {
			return nil, status.Error(codes.Internal, "couldn't edit todo")
		}
	}

	return &todopb.Empty{}, nil
}