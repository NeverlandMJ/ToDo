package server

import (
	"context"
	"time"

	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/google/uuid"
)

// Repository is a main set for database's methods 
type Repository interface {
	CreateTodo(ctx context.Context, td entity.Todo) error
	GetTodo(ctx context.Context, id uuid.UUID) (entity.Todo, error)
	MarkAsDone(ctx context.Context, userID uuid.UUID, todoID uuid.UUID) error
	DeleteTodo(ctx context.Context, id uuid.UUID) error
	GetAllTodos(ctx context.Context, userID uuid.UUID) ([]entity.Todo, error)
	UpdateTodosBody(ctx context.Context, id uuid.UUID, newBody string) error
	UpdateTodosDeadline(ctx context.Context, id uuid.UUID, deadline time.Time) error
	DeleteDoneTodos(ctx context.Context, userID uuid.UUID) error
	DeletePassedDeadline(ctx context.Context, userID uuid.UUID) error
}
