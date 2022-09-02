package service

import (
	"context"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/validate"
	"github.com/NeverlandMJ/ToDo/user-service/server"
	"github.com/google/uuid"
)

// Service holds a type which implements all the methods of Repository.
// It acts as a middleman between grpc server and database server
type Service struct {
	Repo server.Repository
	Otp  validate.Otp
}

// NewService creates a new Service 
func NewService(repo server.Repository) *Service {
	return &Service{
		Repo: repo,
		Otp:  validate.NewOtp(),
	}
}

// CreateUsernameAndPassword generates default user_name and password for user and send it to the user.
// User uses them to sign in. After signing in user can change those default user_name and password.
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

// CreateUser .. 
func (s Service) CreateUser(ctx context.Context, user entity.User) error {
	user = entity.NewUser(user.UserName, user.Password, user.Phone)
	err := s.Repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// GetUser ...
func (s Service) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	user, err := s.Repo.GetUser(ctx, username, password)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// ChangePassword ...
func (s Service) ChangePassword(ctx context.Context, userID uuid.UUID, oldPW, newPW string) error {
	err := s.Repo.ChangePassword(ctx, userID, oldPW, newPW)
	if err != nil {
		return err
	}

	return nil
}

// ChangeUserName ... 
func (s Service) ChangeUserName(ctx context.Context, userID uuid.UUID, newUN string) error {
	err := s.Repo.ChangeUserName(ctx, userID, newUN)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccount ... 
func (s Service) DeleteAccount(ctx context.Context, userID uuid.UUID, password, userName string) error {
	u, err := s.GetUser(ctx, password, userName)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteAccount(ctx, u.ID, u.Password, u.UserName)
	if err != nil {
		return err
	}

	return nil
}
