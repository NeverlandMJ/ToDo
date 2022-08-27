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

func NewServer(cnfg config.Config, path string) (*Server, error) {
	conn, err := database.Connect(cnfg, path)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: conn,
	}, nil
}

func (s Server) deleteUsers() error {
	if _, err := s.db.Exec(`DELETE FROM users`); err != nil {
		return err
	}

	return nil
}

func (s Server) CreateUser(ctx context.Context, user entity.User) error {
	var tempU entity.User
	err := s.db.QueryRowContext(ctx, `SELECT 
	 id, user_name, password, phone_number, created_at, updated_at, is_blocked
	FROM users WHERE phone_number=$1 `, user.Phone).
	Scan(&tempU.ID, &tempU.UserName, &tempU.Password, &tempU.Phone, &tempU.CreatedAt, &tempU.UpdatedAt, &tempU.IsBlocked)

	if err != sql.ErrNoRows{
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

func (s Server) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	var u entity.User
	err := s.db.QueryRowContext(ctx, `
		SELECT id, user_name, password, phone_number, created_at, updated_at, is_blocked
		FROM users WHERE user_name=$1 AND password=$2`, username, password).
	Scan(&u.ID, &u.UserName, &u.Password, &u.Phone, &u.CreatedAt, &u.UpdatedAt, &u.IsBlocked)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, customErr.ERR_USER_NOT_EXIST
		}
		return entity.User{}, err
	}

	return u, nil
}	
