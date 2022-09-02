package service

import (
	"context"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/validate"
	"github.com/NeverlandMJ/ToDo/user-service/server"
	"github.com/google/uuid"
)

type Service struct {
	Repo server.Repository
	Otp  validate.Otp
}

func NewService(repo server.Repository) *Service {
	return &Service{
		Repo: repo,
		Otp:  validate.NewOtp(),
	}
}

func (s Service) CreateUsernameAndPassword(ctx context.Context) (entity.ResponseUser, error) {
	pw, err := validate.GeneratePassword()
	if err != nil {
		return entity.ResponseUser{}, err
	}
	un := validate.GenerateUserName()

	return entity.ResponseUser{
		UserName: un,
		Password: pw,
	}, nil
}

func (s Service) CreateUser(ctx context.Context, user entity.User) error {
	user = entity.NewUser(user.UserName, user.Password, user.Phone)
	err := s.Repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	user, err := s.Repo.GetUser(ctx, username, password)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s Service) ChangePassword(ctx context.Context, userID uuid.UUID, oldPW, newPW string) error {
	err := s.Repo.ChangePassword(ctx, userID, oldPW, newPW)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) ChangeUserName(ctx context.Context, userID uuid.UUID, newUN string) error {
	err := s.Repo.ChangeUserName(ctx, userID, newUN)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteAccount(ctx context.Context, userID uuid.UUID, password, userName string) error {
	err := s.Repo.DeleteAccount(ctx, userID, password, userName)
	if err != nil {
		return err
	}

	return nil
}
