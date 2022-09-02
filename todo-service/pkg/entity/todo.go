package entity

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Deadline  time.Time `json:"deadline"`
	IsDone    bool      `json:"is_done"`
}

func NewTodo(deadline time.Time, body string, userID uuid.UUID) Todo {
	tm := time.Now().UTC().Format(time.UnixDate)
	created, _ := time.Parse(time.UnixDate, tm)

	return Todo{
		ID:        uuid.New(),
		UserID:    userID,
		Body:      body,
		CreatedAt: created,
		Deadline:  deadline,
		IsDone:    false,
	}
}
