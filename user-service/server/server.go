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
// Server holds databse 
type Server struct {
	db *sql.DB
}

// NewServer returns a new Server with working database attached to it.
// If an error occuras while connecting to database, it returns an error
func NewServer(cnfg config.Config, path string) (*Server, error) {
	conn, err := database.Connect(cnfg, path)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: conn,
	}, nil
}

// deleteTodo this functions was written to be used as a cleanUP function inside intigrations tests
func (s Server) deleteUsers() error {
	if _, err := s.db.Exec(`DELETE FROM users`); err != nil {
		return err
	}

	return nil
}

// CheckIfExists checks if the user with given ID actually exists in database.
// It returns true if the todo exists, otherways false
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

// CreateUser inserts a new user into database.
// It returns customErr.ERR_USER_EXIST if the user already exist OR user_name is already taken
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

// GetUser fetches user's data from database by the given username amd password.
// It returns customErr.ERR_USER_NOT_EXIST if the user doesn't exist.
// It returns customErr.ERR_INCORRECT_PASSWORD if the password is incorrect
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

// ChangePassword changes user's password by userID. For that old password is required.
// If old password is incorrect, it returns customErr.ERR_INCORRECT_PASSWORD.
// If user doesn't exist returns customErr.ERR_USER_NOT_EXIST
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

// ChangeUserName changes user_name by userID.
// If user doesn't exist it returns customErr.ERR_USER_NOT_EXIST
// If user_name already taken it returns customErr.ERR_USER_EXIST
func (s Server) ChangeUserName(ctx context.Context, userID uuid.UUID, newUN string) error {
	exist := s.CheckIfExists(ctx, userID)
	if !exist {
		return customErr.ERR_USER_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `UPDATE users SET user_name=$1 WHERE id=$2`, newUN, userID)
	if err != nil {
		return customErr.ERR_USER_EXIST
	}

	return nil
}

// DeleteAccount deletes user's account by userID.
// For authentication purposes password and user_name are asked
func (s Server) DeleteAccount(ctx context.Context, userID uuid.UUID, pw, un string) error  {
	
	_, err := s.db.ExecContext(ctx, `DELETE FROM users WHERE id=$1 AND user_name = $2`, userID, un)
	if err != nil {
		return err
	}
	return nil
}