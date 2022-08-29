package server

import (
	"context"
	"database/sql"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	"github.com/NeverlandMJ/ToDo/user-service/database"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	customErr "github.com/NeverlandMJ/ToDo/user-service/pkg/error"
	"golang.org/x/crypto/bcrypt"
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
	bp, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `INSERT INTO users
		(id, user_name, password, phone_number, created_at, updated_at, is_blocked)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, user.ID, user.UserName, string(bp), user.Phone, user.CreatedAt, user.UpdatedAt, user.IsBlocked)
	if err != nil {
		return customErr.ERR_USER_EXIST
	}
	return nil
}

func (s Server) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	var u entity.User
	err := s.db.QueryRowContext(ctx, `
		SELECT id, user_name, password, phone_number, created_at, updated_at, is_blocked
		FROM users WHERE user_name=$1`, username).
	Scan(&u.ID, &u.UserName, &u.Password, &u.Phone, &u.CreatedAt, &u.UpdatedAt, &u.IsBlocked)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, customErr.ERR_USER_NOT_EXIST
		}
		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return entity.User{}, customErr.ERR_INCORRECT_PASSWORD
	}

	return u, nil
}	
