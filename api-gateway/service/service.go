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
	RegisterUser(ctx context.Context, code entity.ReqSignUp) (entity.RespUser, error)
	SignIn(ctx context.Context, data entity.ReqSignIn) (string, error)
	ChangePassword(ctx context.Context, userID string, new entity.ReqChangePassword) error
	ChangeUserName(ctx context.Context, userID string, new entity.ReqChangeUsername) error
	DeleteAccount(ctx context.Context, userID string, auth entity.ReqSignIn) error
}

type TodoServiceProvider interface {
	CreateTodo(ctx context.Context, td entity.ReqCreateTodo, userID string) (entity.RespTodo, error)
	GetTodoByID(ctx context.Context, userID, todoID string) (entity.RespTodo, error)
	MarkAsDone(ctx context.Context, userID, todoID string) error
	DeleteTodoByID(ctx context.Context, userID, todoID string) error 
	GetAllTodos(ctx context.Context, userID string) (tds []entity.RespTodo, err error)
	UpdateTodosBody(ctx context.Context, userID string, new entity.ReqUpdateBody) error
	UpdateTodosDeadline(ctx context.Context, userID string, new entity.ReqUpdateDeadline) error
	DeleteDoneTodos(ctx context.Context, userID string) error 
	DeletePassedDeadline(ctx context.Context, userID string) error
}

func NewProvider(userServiceURL, todoServiceURL, redisURL string) Provider {
	return Provider{
		UserServiceProvider: NewGRPCClientUser(userServiceURL, redisURL),
		TodoServiceProvider: NewGRPCClientTodo(todoServiceURL),
	}
}
