package server

import (
	"context"
	"database/sql"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	"github.com/NeverlandMJ/ToDo/user-service/database"
	customErr "github.com/NeverlandMJ/ToDo/user-service/pkg/error"
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
	_, err := s.db.QueryContext(ctx, `SELECT * FROM users WHERE phone_number=$1 `, user.Phone)

	if err != nil {
		return customErr.ERR_USER_EXIST
	}

	_, err = s.db.ExecContext(ctx, `INSERT INTO users
		(id, user_name, password, phone_number, created_at, updated_at, is_blocked)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, user.ID, user.UserName, user.Password, user.Phone, user.CreatedAt, user.UpdatedAt, user.IsBlocked)
	if err != nil {
		return  err
	}
	return nil
}

