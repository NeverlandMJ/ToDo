package service

import (
	"context"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/validate"
	"github.com/NeverlandMJ/ToDo/user-service/server"
)

type Service struct {
	Repo server.Repository
	Otp validate.Otp
}

func NewService(repo server.Repository) *Service {
	return &Service{
		Repo: repo,
		Otp: validate.NewOtp(),
	}
}


func (s Service) CreateUserNameAndPassword(ctx context.Context) (entity.ResponseUser, error)  {
	pw, err := validate.GeneratePassword()
	if err != nil {
		return entity.ResponseUser{}, err
	}
	un := validate.	GenerateUserName()

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
