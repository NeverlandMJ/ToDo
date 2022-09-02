package service

import (
	"context"
	"time"

	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/todo-service/server"
	"github.com/google/uuid"
)

type Service struct {
	Repo server.Repository
}

func NewService(repo server.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s Service) CreateTodo(ctx context.Context, td entity.Todo) (entity.Todo, error) {
	newTd := entity.NewTodo(td.Deadline, td.Body, td.UserID)

	err := s.Repo.CreateTodo(ctx, newTd)
	if err != nil {
		return entity.Todo{}, err
	}

	return newTd, nil
}

func (s Service) GetTodo(ctx context.Context, id uuid.UUID) (entity.Todo, error) {
	got, err := s.Repo.GetTodo(ctx, id)
	if err != nil {
		return entity.Todo{}, err
	}
	return got, nil
}

func (s Service) MarkAsDone(ctx context.Context, id uuid.UUID) error {
	err := s.Repo.MarkAsDone(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) DeleteTodoByID(ctx context.Context, id uuid.UUID) error  {
	err := s.Repo.DeleteTodo(ctx, id)
	if err != nil {
		return  err
	}

	return nil
}

func (s Service) GetAllTodos(ctx context.Context, userID uuid.UUID) ([]entity.Todo, error)  {
	tds, err := s.Repo.GetAllTodos(ctx, userID)
	if err != nil {
		return []entity.Todo{}, err
	}

	return tds, nil
}

func (s Service) UpdateTodosBody(ctx context.Context, id uuid.UUID, newBody string) error {
	err := s.Repo.UpdateTodosBody(ctx, id, newBody)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateTodosDeadline(ctx context.Context, id uuid.UUID, newDeadline time.Time) error  {
	err := s.Repo.UpdateTodosDeadline(ctx, id, newDeadline)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteDoneTodos(ctx context.Context, userID uuid.UUID) error  {
	err := s.Repo.DeleteDoneTodos(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeletePassedDeadline(ctx context.Context, userID uuid.UUID) error {
	err := s.Repo.DeletePassedDeadline(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}