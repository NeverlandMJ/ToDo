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

// gRPCServer holds grpc TodoService server and service of the project
type gRPCServer struct {
	todopb.UnimplementedTodoServiceServer
	svc service.Service
}

// NewgRPCServer returns a new grpc server with the given service attached 
func NewgRPCServer(svc service.Service) *gRPCServer {
	return &gRPCServer{
		svc: svc,
	}
}

// CreateTodo creates a new todo with given cridentials 
func (g *gRPCServer) CreateTodo(ctx context.Context, req *todopb.RequestTodo) (*todopb.ResponseTodo, error) {
	id, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "UserID is not uuid")
	}

	ddl, err := time.Parse(time.UnixDate, req.Deadline)
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

// GetTodoByID gets a todo by given ID
func (g *gRPCServer) GetTodoByID(ctx context.Context, req *todopb.RequestGetTodo) (*todopb.ResponseTodo, error) {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "userID is not uuid")
	}
	todoID, err := uuid.Parse(req.GetTodoId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "todoID is not uuid")
	}

	td, err := g.svc.GetTodo(ctx, userID, todoID)
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

// MarkAsDone marks todo as done by the given todo ID
func (g *gRPCServer) MarkAsDone(ctx context.Context, req *todopb.RequestMarkAsDone) (*todopb.Empty, error) {
 	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "userID is not uuid")
	}
 	todoID, err := uuid.Parse(req.GetTodoId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "todoID is not uuid")
	}

	err = g.svc.MarkAsDone(ctx, userID, todoID)
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

// DeleteTodoByID deletes todo by the given ID
func (g *gRPCServer) DeleteTodoByID(ctx context.Context, req *todopb.RequestDeleteTodo) (*todopb.Empty, error)  {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "userID is not uuid")
	}
 	todoID, err := uuid.Parse(req.GetTodoId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "todoID is not uuid")
	}

	err = g.svc.DeleteTodoByID(ctx, userID, todoID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, customerr.ERR_TODO_NOT_EXIST) {
			return nil, status.Error(codes.NotFound, "todo with given id is not found")
		}else {
			return nil, status.Error(codes.Internal, "couldn't delete todo")
		}
	}
	return &todopb.Empty{}, nil
}

// GetAllTodos gets all todos by the user ID 
func (g *gRPCServer) GetAllTodos(ctx context.Context, req *todopb.RequestUserID) (*todopb.ResponseAllTodos, error)  {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}
	tds, err := g.svc.GetAllTodos(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't fetch todos")
	}

	protoTds := make([]*todopb.ResponseTodo, 0, len(tds))
	for _, td := range tds {
		protoTd := todopb.ResponseTodo{
			Id: td.ID.String(),
			UserID: td.UserID.String(),
			Body: td.Body,
			CreatedAt: td.CreatedAt.String(),
			Deadline: td.Deadline.String(),
			IsDone: td.IsDone,
		}
		protoTds =append(protoTds, &protoTd)
	}

	return &todopb.ResponseAllTodos{
		ResponseTodos: protoTds,
	}, nil
}

// UpdateTodosBody updates todo's body by the ID 
func (g *gRPCServer) UpdateTodosBody(ctx context.Context, req *todopb.RequestUpdateTodosBody) (*todopb.Empty, error)  {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "userID is not uuid")
	}
 	todoID, err := uuid.Parse(req.GetTodoId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "todoID is not uuid")
	}

	err = g.svc.UpdateTodosBody(ctx, userID, todoID, req.NewBody)
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

// UpdateTodosDeadline updates todo's deadline by the ID
func (g *gRPCServer) UpdateTodosDeadline(ctx context.Context, req *todopb.RequestUpdateTodosDeadline) (*todopb.Empty, error) {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "userID is not uuid")
	}
 	todoID, err := uuid.Parse(req.GetTodoId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "todoID is not uuid")
	}


	tm, err := time.Parse(time.UnixDate, req.NewDeadline)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "given time is not parsable")
	}

	err = g.svc.UpdateTodosDeadline(ctx, userID, todoID, tm)	
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

// DeleteDoneTodos deletes all todos by userID which was marked as done
func (g *gRPCServer) DeleteDoneTodos(ctx context.Context, req *todopb.RequestUserID) (*todopb.Empty, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	err = g.svc.DeleteDoneTodos(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "couldn't delete todo")
	}

	return &todopb.Empty{}, nil
}

// DeletePassedDeadline deletes all todos whichs deadline was passed 
func (g *gRPCServer) DeletePassedDeadline(ctx context.Context, req *todopb.RequestUserID) (*todopb.Empty, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "ID is not uuid")
	}

	err = g.svc.DeletePassedDeadline(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "couldn't delete todo")
	}

	return &todopb.Empty{}, nil
}