package server

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/NeverlandMJ/ToDo/todo-service/config"
	customerr "github.com/NeverlandMJ/ToDo/todo-service/pkg/customERR"
	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var testDeadline = time.Date(2022, time.September, 8, 7, 0, 0, 0, time.UTC)
var testUserID = uuid.New()

func TestServer_CreateTodo(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		err := insertTestUser(s)
		require.NoError(t, err)

		todo, err := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		require.NoError(t, err)

		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)
	})
}

func TestServer_GetTodo(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		todo, err := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		require.NoError(t, err)

		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		td, err := s.GetTodo(context.Background(), todo.ID)
		

		require.NoError(t, err)
		require.Equal(t, todo, td)

	})
	t.Run("todo doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		td, err := s.GetTodo(context.Background(), uuid.New())

		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
		require.Equal(t, entity.Todo{}, td)

	})

}

func TestServer_MarkAsDone(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		todo, err := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		require.NoError(t, err)

		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		err = s.MarkAsDone(context.Background(), todo.ID)
		require.NoError(t, err)

		td, err := s.GetTodo(context.Background(), todo.ID)

		require.NoError(t, err)
		require.Equal(t, true, td.IsDone)

	})

	t.Run("todo doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		err = s.MarkAsDone(context.Background(), uuid.New())
		
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
}

func newServer(t *testing.T) *Server {
	t.Helper()
	serv, err := NewServer(
		config.Config{
			Host:             "localhost",
			Port:             "8080",
			PostgresHost:     "localhost",
			PostgresPort:     "5432",
			PostgresUser:     "sunbula",
			PostgresPassword: "2307",
			PostgresDB:       "todo_test",
		},
		"file://./../migrations/",
	)

	require.NoError(t, err)

	return serv
}

func insertTestUser(s *Server) error {
	_, err := s.db.Exec(`
		INSERT INTO users
		(id, user_name, password, phone_number)
		VALUES ($1, $2, $3, $4)
	`, testUserID, "sunbula", "123", "+998887882307")
	return err
}

func cleanUpFn(s *Server) func() {
	return func() {
		if err := s.deleteTodo(); err != nil {
			log.Println("CLEANUP OF DB FAILED!")
		}
	}
}
