package service

import (
	"context"
	"fmt"
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
		fmt.Println("Todo error", err)
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

func (c todoServiceGRPCClient) DeleteTodoByID(ctx context.Context, userID, todoID string) error {
	_, err := c.client.DeleteTodoByID(ctx, &todopb.RequestDeleteTodo{
		UserId: userID,
		TodoId: todoID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c todoServiceGRPCClient) GetAllTodos(ctx context.Context, userID string) (tds []entity.RespTodo, err error) {
	resp, err := c.client.GetAllTodos(ctx, &todopb.RequestUserID{
		Id: userID,
	})

	if err != nil {
		return nil, err
	}
	
	for _, grpcTD := range resp.ResponseTodos {
		td := entity.RespTodo{}
		
		td.ID = grpcTD.Id
		td.UserID = grpcTD.UserID
		td.Body = grpcTD.Body
		td.CreatedAt = grpcTD.CreatedAt
		td.Deadline = grpcTD.Deadline
		td.IsDone = grpcTD.IsDone

		tds = append(tds, td)
	}
	
	return tds, nil
}

func (c todoServiceGRPCClient) UpdateTodosBody(ctx context.Context, userID string, new entity.ReqUpdateBody) error {
	_, err := c.client.UpdateTodosBody(ctx, &todopb.RequestUpdateTodosBody{
		UserId: userID,
		TodoId: new.TodoID,
		NewBody: new.Body,
	})
	if err != nil {
		return  err
	}

	return nil
}

func (c todoServiceGRPCClient) UpdateTodosDeadline(ctx context.Context, userID string, new entity.ReqUpdateDeadline) error {
	_, err := c.client.UpdateTodosDeadline(ctx, &todopb.RequestUpdateTodosDeadline{
		TodoId: new.TodoID,
		UserId: userID,
		NewDeadline: new.Deadline,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c todoServiceGRPCClient) DeleteDoneTodos(ctx context.Context, userID string) error {
	_, err := c.client.DeleteDoneTodos(ctx, &todopb.RequestUserID{
		Id: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c todoServiceGRPCClient) DeletePassedDeadline(ctx context.Context, userID string) error {
	_, err := c.client.DeletePassedDeadline(ctx, &todopb.RequestUserID{
		Id: userID,
	})
	if err != nil {
		return err
	}
	return nil
}