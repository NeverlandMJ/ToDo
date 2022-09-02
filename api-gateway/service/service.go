package service

import (
	"context"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/entity"
)

type Provider struct {
	UserServiceProvider
	TodoServiceProvider
}

type UserServiceProvider interface {
	SendCode(ctx context.Context, ph entity.ReqPhone) (entity.ReqPhone, error)
	RegisterUser(ctx context.Context, code entity.ReqCode) (entity.RespUser, error)
	SignIn(ctx context.Context, data entity.ReqSignIn) (string, error)
}

type TodoServiceProvider interface {
	CreateTodo(ctx context.Context, td entity.ReqCreateTodo, userID string) (entity.RespTodo, error)
	GetTodoByID(ctx context.Context, userID, todoID string) (entity.RespTodo, error)
	MarkAsDone(ctx context.Context, userID, todoID string) error
}

func NewProvider(userServiceURL, todoServiceURL string) Provider {
	return Provider{
		UserServiceProvider: NewGRPCClientUser(userServiceURL),
		TodoServiceProvider: NewGRPCClientTodo(todoServiceURL),
	}
}


