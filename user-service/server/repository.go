package server

import (
	"context"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/google/uuid"
)

// Repository is a main set for database's methods 
type Repository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, username, password string) (entity.User, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, oldPW, newPW string) error
	ChangeUserName(ctx context.Context, userID uuid.UUID, newUN string) error
	DeleteAccount(ctx context.Context, userID uuid.UUID, pw, un  string) error
}
