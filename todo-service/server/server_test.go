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
var newDeadline = time.Date(2022, time.September, 13, 6, 30, 0, 0, time.UTC)
var passedDeadlie = time.Date(2022, time.July, 23, 18, 0, 0, 0, time.UTC)

func TestServer_CreateTodo(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		err := insertTestUser(s)
		require.NoError(t, err)

		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)

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

		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)

		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		td, err := s.GetTodo(context.Background(), todo.ID)
		td.CreatedAt = todo.CreatedAt

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

		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)

		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		err = s.MarkAsDone(context.Background(), todo.UserID, todo.ID)
		require.NoError(t, err)

		td, err := s.GetTodo(context.Background(), todo.ID)

		require.NoError(t, err)
		require.Equal(t, true, td.IsDone)

	})

	t.Run("todo doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		err = s.MarkAsDone(context.Background(), uuid.New(), uuid.New())
		
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
}

func TestServer_DeleteTodo(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)

		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		err = s.DeleteTodo(context.Background(), todo.ID)
		require.NoError(t, err)

		td, err := s.GetTodo(context.Background(), todo.ID)
		
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
		require.EqualValues(t, entity.Todo{}, td)
	})

	t.Run("todo doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		err = s.DeleteTodo(context.Background(), uuid.New())
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)

	})
}

func TestServer_GetAllTodos(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		want := []entity.Todo{}

		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)
		want = append(want, todo)

		todo = entity.NewTodo(testDeadline, "bake a cake", testUserID)
		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)
		want = append(want, todo)


		todo = entity.NewTodo(testDeadline, "get up early", testUserID)
		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)
		want = append(want, todo)

		got, err :=s.GetAllTodos(context.Background(), testUserID)
		require.NoError(t, err)
		require.EqualValues(t, want, got)		
	})
}

func TestServer_UpdateTodosBody(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)
		
		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		err = s.UpdateTodosBody(context.Background(), todo.ID, "make a cake")
		require.NoError(t, err)

		got, err := s.GetTodo(context.Background(), todo.ID)
		todo.Body = "make a cake"

		require.NoError(t, err)
		require.EqualValues(t, todo, got)
	})

	t.Run("todo doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		err = s.UpdateTodosBody(context.Background(), uuid.New(), "make a cake")
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
}

func TestServer_UpdateTodosDeadline(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)
		
		todo := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		err = s.CreateTodo(context.Background(), todo)
		require.NoError(t, err)

		err = s.UpdateTodosDeadline(context.Background(), todo.ID, newDeadline)
		require.NoError(t, err)

		got, err := s.GetTodo(context.Background(), todo.ID)
		require.NoError(t, err)

		require.EqualValues(t, newDeadline, got.Deadline)
	})
	t.Run("todo doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		err = s.UpdateTodosDeadline(context.Background(), uuid.New(), newDeadline)
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
}

func TestServer_DeletDoneTodos(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		want := []entity.Todo{}

		todo1 := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		err = s.CreateTodo(context.Background(), todo1)
		require.NoError(t, err)
		want = append(want, todo1)

		todo2 := entity.NewTodo(testDeadline, "bake a cake", testUserID)
		err = s.CreateTodo(context.Background(), todo2)
		require.NoError(t, err)
		want = append(want, todo2)


		todo3 := entity.NewTodo(testDeadline, "get up early", testUserID)
		err = s.CreateTodo(context.Background(), todo3)
		require.NoError(t, err)

		err = s.MarkAsDone(context.Background(), todo3.UserID, todo3.ID)
		require.NoError(t, err)

		err = s.DeleteDoneTodos(context.Background(), testUserID)
		require.NoError(t, err)
	
		got, err := s.GetAllTodos(context.Background(), testUserID)
		require.NoError(t, err)
		require.EqualValues(t, want, got)		
	})
}

func TestServer_DeletDeadlinePassed(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("should pass", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		err := insertTestUser(s)
		require.NoError(t, err)

		want := []entity.Todo{}

		todo1 := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		err = s.CreateTodo(context.Background(), todo1)
		require.NoError(t, err)
		want = append(want, todo1)

		todo2 := entity.NewTodo(testDeadline, "bake a cake", testUserID)
		err = s.CreateTodo(context.Background(), todo2)
		require.NoError(t, err)
		want = append(want, todo2)

		todo3 := entity.NewTodo(passedDeadlie, "get up early", testUserID)
		err = s.CreateTodo(context.Background(), todo3)
		require.NoError(t, err)
		
		err = s.DeletePassedDeadline(context.Background(), testUserID)
		require.NoError(t, err)
	
		got, err := s.GetAllTodos(context.Background(), testUserID)
		require.NoError(t, err)
		require.EqualValues(t, want, got)	
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
