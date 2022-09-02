package service

import (
	"context"
	"time"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/entity"
	"github.com/NeverlandMJ/ToDo/api-gateway/v1/todopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type todoServiceGRPCClient struct {
	client todopb.TodoServiceClient
}

func NewGRPCClientTodo(url string) todoServiceGRPCClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	client := todopb.NewTodoServiceClient(conn)

	return todoServiceGRPCClient{
		client: client,
	}
}

func (c todoServiceGRPCClient) CreateTodo(ctx context.Context, td entity.ReqCreateTodo, userID string) (entity.RespTodo, error) {
	resp, err := c.client.CreateTodo(ctx, &todopb.RequestTodo{
		UserId: userID,
		Body: td.Body,
		Deadline: td.Deadline,
	})

	if err != nil {
		return entity.RespTodo{}, err
	}

	return entity.RespTodo{
		ID: resp.Id,
		UserID: resp.UserID,
		Body: resp.Body,
		CreatedAt: resp.CreatedAt,
		Deadline: resp.Deadline,
		IsDone: resp.IsDone,
	}, nil
}

func (c todoServiceGRPCClient) MarkAsDone(ctx context.Context, userID, todoID string) error {
	_, err := c.client.MarkAsDone(ctx, &todopb.RequestMarkAsDone{
		UserId: userID,
		TodoId: todoID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c todoServiceGRPCClient) GetTodoByID(ctx context.Context, userID, todoID string) (entity.RespTodo, error) {
	resp, err := c.client.GetTodoByID(ctx, &todopb.RequestGetTodo{
		UserId: userID,
		TodoId: todoID,
	})

	if err != nil {
		return entity.RespTodo{} ,err
	}

	return entity.RespTodo{
		ID: resp.Id,
		UserID: resp.UserID,
		Body: resp.Body,
		CreatedAt: resp.CreatedAt,
		Deadline: resp.Deadline,
		IsDone: resp.IsDone,
	}, nil
}