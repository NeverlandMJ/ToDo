package server

import (
	"context"

	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/google/uuid"
)

type Repository interface {
	CreateTodo(ctx context.Context, td entity.Todo) error
	GetTodo(ctx context.Context, id uuid.UUID) (entity.Todo, error)
	MarkAsDone(ctx context.Context, id uuid.UUID) error
}