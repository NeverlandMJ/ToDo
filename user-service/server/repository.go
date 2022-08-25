package server

import (
	"context"

	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, user entity.User) error
}
