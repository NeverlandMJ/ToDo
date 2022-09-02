package service

import (
	"context"
	"time"

	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/todo-service/server"
	"github.com/google/uuid"
)

// Service holds a type which implements all the methods of Repository.
// It acts as a middleman between grpc server and database server
type Service struct {
	Repo server.Repository
}

// NewService creates a new Service 
func NewService(repo server.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// CreateTodo ....
func (s Service) CreateTodo(ctx context.Context, td entity.Todo) (entity.Todo, error) {
	newTd := entity.NewTodo(td.Deadline, td.Body, td.UserID)

	err := s.Repo.CreateTodo(ctx, newTd)
	if err != nil {
		return entity.Todo{}, err
	}

	return newTd, nil
}

// GetTodo ...
func (s Service) GetTodo(ctx context.Context, userID uuid.UUID, todoID uuid.UUID) (entity.Todo, error) {
	got, err := s.Repo.GetTodo(ctx, userID, todoID)
	if err != nil {
		return entity.Todo{}, err
	}
	return got, nil
}

// MarkAsDone ...
func (s Service) MarkAsDone(ctx context.Context, userID uuid.UUID, todoID uuid.UUID) error {
	err := s.Repo.MarkAsDone(ctx, userID, todoID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodoByID ...
func (s Service) DeleteTodoByID(ctx context.Context, userID, todoID uuid.UUID) error  {
	err := s.Repo.DeleteTodo(ctx, userID, todoID)
	if err != nil {
		return  err
	}

	return nil
}

// GetAllTodos ...
func (s Service) GetAllTodos(ctx context.Context, userID uuid.UUID) ([]entity.Todo, error)  {
	tds, err := s.Repo.GetAllTodos(ctx, userID)
	if err != nil {
		return []entity.Todo{}, err
	}

	return tds, nil
}

// UpdateTodosBody ...
func (s Service) UpdateTodosBody(ctx context.Context, userID, todoID uuid.UUID, newBody string) error {
	err := s.Repo.UpdateTodosBody(ctx, userID, todoID, newBody)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTodosDeadline ...
func (s Service) UpdateTodosDeadline(ctx context.Context, userID, todoID uuid.UUID, newDeadline time.Time) error  {
	err := s.Repo.UpdateTodosDeadline(ctx, userID, todoID, newDeadline)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDoneTodos ...
func (s Service) DeleteDoneTodos(ctx context.Context, userID uuid.UUID) error  {
	err := s.Repo.DeleteDoneTodos(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

// DeletePassedDeadline ...
func (s Service) DeletePassedDeadline(ctx context.Context, userID uuid.UUID) error {
	err := s.Repo.DeletePassedDeadline(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}