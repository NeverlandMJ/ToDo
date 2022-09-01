package entity

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Deadline  time.Time `json:"deadline"`
	IsDone    bool      `json:"is_done"`
}

func NewTodo(deadline time.Time, body string, userID uuid.UUID) (Todo, error) {
	// tm := time.Now().Format("UnixDate")
	// created, err := time.Parse("UnixDate", tm)
	// if err != nil {
	// 	return Todo{}, err
	// }

	return Todo{
		ID:        uuid.New(),
		UserID:    userID,
		Body:      body,
		CreatedAt: time.Now().UTC(),
		Deadline:  deadline,
		IsDone:    false,
	}, nil
}
