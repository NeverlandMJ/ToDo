package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NeverlandMJ/ToDo/todo-service/config"
	"github.com/NeverlandMJ/ToDo/todo-service/database"
	customerr "github.com/NeverlandMJ/ToDo/todo-service/pkg/customERR"
	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/google/uuid"
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

func (s Server) CreateTodo(ctx context.Context, td entity.Todo) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO todos 
		(id, user_id, body, created_at, deadline, is_done)
		VALUES
		($1, $2, $3, $4, $5, $6)
	`, td.ID, td.UserID, td.Body, td.CreatedAt, td.Deadline, td.IsDone)

	if err != nil {
		return err
	}

	return nil
}

func (s Server) GetTodo(ctx context.Context, id uuid.UUID) (entity.Todo, error) {
	var td entity.Todo
	exist := s.CheckIfExists(ctx, id)
	if !exist {
		return td, customerr.ERR_TODO_NOT_EXIST
	}
	err := s.db.QueryRowContext(ctx, `
		SELECT * FROM todos WHERE id=$1
	`, id).Scan(&td.ID, &td.UserID, &td.Body, &td.CreatedAt, &td.Deadline, &td.IsDone)

	if err != nil {
		return td, err
	}

	return td, nil
}

func (s Server) MarkAsDone(ctx context.Context, id uuid.UUID) error {
	exist := s.CheckIfExists(ctx, id)
	if !exist {
		return customerr.ERR_TODO_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `
		UPDATE todos SET is_done = true WHERE id=$1 
	`, id)

	if err != nil {
		return err
	}

	return nil
}

func (s Server) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	exists := s.CheckIfExists(ctx, id)
	if !exists {
		return customerr.ERR_TODO_NOT_EXIST
	}
	_, err := s.db.ExecContext(ctx, `DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) GetAllTodos(ctx context.Context, userID uuid.UUID) ([]entity.Todo, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT * FROM todos WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	tds := []entity.Todo{}

	for rows.Next() {
		td := entity.Todo{}
		err := rows.Scan(&td.ID, &td.UserID, &td.Body, &td.CreatedAt, &td.Deadline, &td.IsDone)
		if err != nil {
			return nil, err
		}
		tds = append(tds, td)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return tds, nil
}

func (s Server) UpdateTodosBody(ctx context.Context, id uuid.UUID, newBody string) error {
	exists := s.CheckIfExists(ctx, id)
	if !exists {
		return customerr.ERR_TODO_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `UPDATE todos SET body=$2 WHERE id=$1`, id, newBody)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) UpdateTodosDeadline(ctx context.Context, id uuid.UUID, deadline time.Time) error {
	exists := s.CheckIfExists(ctx, id)
	if !exists {
		return customerr.ERR_TODO_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `UPDATE todos SET deadline=$1 WHERE id=$2`, deadline, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Server)  DeleteDoneTodos(ctx context.Context, userID uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM todos WHERE user_id=$1 AND is_done=true`, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) DeletePassedDeadline(ctx context.Context, userID uuid.UUID) error {
	tm := time.Now().UTC().Format(time.UnixDate)
	now, _ := time.Parse(time.UnixDate, tm)

	_, err := s.db.ExecContext(ctx, `DELETE FROM todos WHERE user_id=$1 AND deadline < $2`, userID, now)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) CheckIfExists(ctx context.Context, id uuid.UUID) bool {
	var exist bool
	err := s.db.QueryRowContext(ctx, `
		select exists(select 1 from todos where id=$1)
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

func (s Server) deleteTodo() error {
	_, err := s.db.Exec(`
		DELETE FROM todos
	`)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(`
		DELETE FROM users
	`)
	if err != nil {
		return err
	}

	return nil
}
