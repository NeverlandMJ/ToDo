package server

import (
	"context"
	"database/sql"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	"github.com/NeverlandMJ/ToDo/user-service/database"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
)

type Server struct {
	db *sql.DB
}

func NewServer(cnfg config.Config) (*Server, error) {
	conn, err := database.Connect(cnfg)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: conn,
	}, nil
}

func (s Server) CreateUser(ctx context.Context, user entity.User) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO users
		(id, user_name, password, phone_number, created_at, updated_at, is_blocked)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return  err
	}
	return nil
}
