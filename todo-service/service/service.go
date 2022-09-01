package service

import (
	"context"

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
	newTd, err := entity.NewTodo(td.Deadline, td.Body, td.UserID)
	if err != nil {
		return entity.Todo{}, err
	}

	err = s.Repo.CreateTodo(ctx, newTd)
	if err != nil {
		return entity.Todo{}, err
	}

	return newTd, nil
}

func (s Service) GetTodo(ctx context.Context, id uuid.UUID) (entity.Todo, error) {
	gotTd, err := s.Repo.GetTodo(ctx, id)
	if err != nil {
		return entity.Todo{}, err
	}
	return gotTd, nil
}

func (s Service) MarkAsDone(ctx context.Context, id uuid.UUID) error {
	err := s.Repo.MarkAsDone(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
