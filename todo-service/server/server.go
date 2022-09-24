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

// Server holds databse 
type Server struct {
	db *sql.DB
}

// NewServer returns a new Server with working database attached to it.
// If an error occuras while connecting to database, it returns an error
func NewServer(cnfg config.Config) (*Server, error) {
	conn, err := database.Connect(cnfg)
	if err != nil {
		return nil, err
	}
	return &Server{
		db: conn,
	}, nil
}

// CreateTodo inserts todo into database
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

// GetTodo fetches todo by id. If there is not todo by the given ID, it returns customerr.ERR_TODO_NOT_EXIST
func (s Server) GetTodo(ctx context.Context, userID uuid.UUID, todoID uuid.UUID) (entity.Todo, error) {
	var td entity.Todo
	exist := s.CheckIfExists(ctx, todoID)
	if !exist {
		return td, customerr.ERR_TODO_NOT_EXIST
	}
	err := s.db.QueryRowContext(ctx, `
		SELECT * FROM todos WHERE id=$1 AND user_id=$2
	`, todoID, userID).Scan(&td.ID, &td.UserID, &td.Body, &td.CreatedAt, &td.Deadline, &td.IsDone)

	if err != nil {
		return td, err
	}

	return td, nil
}

// MarkAsDone updates todo's is_done field by changing it to true. 
// If there is not todo by the given ID, it returns customerr.ERR_TODO_NOT_EXIST
func (s Server) MarkAsDone(ctx context.Context, userID uuid.UUID, todoID uuid.UUID) error {
	exist := s.CheckIfExists(ctx, todoID)
	if !exist {
		return customerr.ERR_TODO_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `
		UPDATE todos SET is_done = true WHERE id=$1 AND user_id=$2
	`, todoID, userID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteTodo deletes todo from database by the given ID.
// If there is not todo by the given ID, it returns customerr.ERR_TODO_NOT_EXIST
func (s Server) DeleteTodo(ctx context.Context, userID uuid.UUID, todoID uuid.UUID) error {
	exists := s.CheckIfExists(ctx, todoID)
	if !exists {
		return customerr.ERR_TODO_NOT_EXIST
	}
	_, err := s.db.ExecContext(ctx, `DELETE FROM todos WHERE id=$1 AND user_id = $2`, todoID, userID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllTodos fetches all todos by userID
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

// UpdateTodosBody updates todo's body.
// If there is not todo by the given ID, it returns customerr.ERR_TODO_NOT_EXIST
func (s Server) UpdateTodosBody(ctx context.Context, userID uuid.UUID, todoID uuid.UUID, newBody string) error {
	exists := s.CheckIfExists(ctx, todoID)
	if !exists {
		return customerr.ERR_TODO_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `UPDATE todos SET body=$1 WHERE id=$2 AND user_id=$3`, newBody, todoID, userID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTodosDeadline updates todo's deadline.
// If there is not todo by the given ID, it returns customerr.ERR_TODO_NOT_EXIST
func (s Server) UpdateTodosDeadline(ctx context.Context, userID uuid.UUID, todoID uuid.UUID, deadline time.Time) error {
	exists := s.CheckIfExists(ctx, todoID)
	if !exists {
		return customerr.ERR_TODO_NOT_EXIST
	}

	_, err := s.db.ExecContext(ctx, `UPDATE todos SET deadline=$1 WHERE id=$2 AND user_id=$3`, deadline, todoID, userID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDoneTodos deletes done todos by the userID
func (s Server)  DeleteDoneTodos(ctx context.Context, userID uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM todos WHERE user_id=$1 AND is_done=true`, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeletePassedDeadline deletes todos which's deadline have already passed
func (s Server) DeletePassedDeadline(ctx context.Context, userID uuid.UUID) error {
	tm := time.Now().UTC().Format(time.UnixDate)
	now, _ := time.Parse(time.UnixDate, tm)

	_, err := s.db.ExecContext(ctx, `DELETE FROM todos WHERE user_id=$1 AND deadline < $2`, userID, now)
	if err != nil {
		return err
	}

	return nil
}

// CheckIfExists checks if the todo with given ID actually exists in database.
// It returns true if the todo exists, otherways false
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

// deleteTodo this functions was written to be used as a cleanUP function inside intigrations tests
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
