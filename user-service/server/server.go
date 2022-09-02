package server

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	"github.com/NeverlandMJ/ToDo/user-service/database"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/customErr"
	"github.com/google/uuid"
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

func (s Server) CheckIfExists(ctx context.Context, id uuid.UUID) bool {
	var exist bool
	err := s.db.QueryRowContext(ctx, `
		select exists(select 1 from users where id=$1)
	`, id).Scan(&exist)

	if err != nil {
		fmt.Println(err)
		return false
	}
	if !exist {
		return false
	} else {
		return true
	}
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
		} else {
			return entity.User{}, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return entity.User{}, customErr.ERR_INCORRECT_PASSWORD
	}

	return u, nil
}

func (s Server) ChangePassword(ctx context.Context, userID uuid.UUID, oldPW, newPW string) error {
	exist := s.CheckIfExists(ctx, userID)
	if !exist {
		return customErr.ERR_USER_NOT_EXIST
	}
	
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Commit()

	var actual string
	err = tx.QueryRowContext(ctx, `SELECT password FROM users WHERE id=$1`, userID ).Scan(&actual)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(actual), []byte(oldPW))
	if err != nil {
		tx.Rollback()
		return customErr.ERR_INCORRECT_PASSWORD
	} 

	newHashed, err := bcrypt.GenerateFromPassword([]byte(newPW), bcrypt.MinCost)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE users SET password=$1 WHERE id=$2`, newHashed, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	return nil
}

func (s Server) ChangeUserName(ctx context.Context, userID uuid.UUID, newUN string) error {
	exist := s.CheckIfExists(ctx, userID)
	if !exist {
		return customErr.ERR_USER_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `UPDATE users SET user_name=$1 WHERE id=$2`, newUN, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) DeleteAccount(ctx context.Context, userID uuid.UUID, pw, un string) error  {
	exist := s.CheckIfExists(ctx, userID)
	if !exist {
		return customErr.ERR_USER_NOT_EXIST
	}

	u, err := s.GetUser(ctx, un, pw)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `DELETE FROM users WHERE id=$1 AND user_name = $2`, userID, u.UserName)
	if err != nil {
		return err
	}
	return nil
}