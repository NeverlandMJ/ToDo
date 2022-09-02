package entity

import (
	"time"

	"github.com/google/uuid"
)

// User holds crideantials of user
type User struct {
	ID            uuid.UUID    `json:"id"`
	UserName      string    `json:"user_name"`
	Password      string    `json:"password"`
	Phone         string    `json:"phone"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsBlocked     bool      `json:"is_blocked"`
}

// RequestUser used to send TOTP code
type RequestUser struct {
	Phone string `json:"email"`
}
// ResponseUser used to generate default user name and password
type ResponseUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// NewUser creates a new user
func NewUser(username, password, phone string) User {
	tm := time.Now().UTC().Format(time.UnixDate)
	t, _ := time.Parse(time.UnixDate, tm)
	
	return User{
		ID: uuid.New() ,
		UserName: username,
		Password: password,
		Phone: phone,
		CreatedAt: t,
		UpdatedAt: t,
		IsBlocked: false,
	}
}