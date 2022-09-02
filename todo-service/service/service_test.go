package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/NeverlandMJ/ToDo/todo-service/mocks"
	customerr "github.com/NeverlandMJ/ToDo/todo-service/pkg/customERR"
	"github.com/NeverlandMJ/ToDo/todo-service/pkg/entity"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var testDeadline = time.Date(2022, time.September, 8, 7, 0, 0, 0, time.UTC)
var testUserID = uuid.New()
var testBody = "get up early"

func newRepos(t *testing.T) *mocks.MockRepository {
	ctl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(ctl)
	return mockRepo
}

func TestService_CreateTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		want := entity.NewTodo(testDeadline, testBody, testUserID)
		repo := newRepos(t)
		repo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return( nil)

		s := NewService(repo)
		got, err := s.CreateTodo(context.Background(), want)

		require.NoError(t, err)
		require.EqualValues(t, want.Deadline, got.Deadline)
		require.EqualValues(t, want.Body, got.Body)
		require.EqualValues(t, want.UserID, got.UserID)

	})
	t.Run("should return error", func(t *testing.T) {
		td := entity.NewTodo(testDeadline, testBody, testUserID)
		repo := newRepos(t)
		repo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

		s := NewService(repo)
		got, err := s.CreateTodo(context.Background(), td)
		require.Error(t, err)
		require.Equal(t, entity.Todo{}, got)

	})

}

func TestService_GetTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		want := entity.NewTodo(testDeadline, testBody, testUserID)
		
		repo := newRepos(t)
		repo.EXPECT().GetTodo(gomock.Any(), gomock.Any()).Return(want, nil)

		s := NewService(repo)
		got, err := s.GetTodo(context.Background(), uuid.New())
		
		require.NoError(t, err)
		require.EqualValues(t, want, got)
	})

	t.Run("should return error", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().GetTodo(gomock.Any(), gomock.Any()).Return(entity.Todo{}, fmt.Errorf("error"))

		s := NewService(repo)
		got, err := s.GetTodo(context.Background(), uuid.New())
		
		require.Error(t, err)
		require.EqualValues(t, entity.Todo{}, got)
	})
}

func TestService_MarkAsDone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().MarkAsDone(gomock.Any(), gomock.Any()).Return(nil)

		s := NewService(repos)
		err := s.MarkAsDone(context.Background(), uuid.New())
		require.NoError(t, err)
	})

	t.Run("returns error", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().MarkAsDone(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

		s := NewService(repos)
		err := s.MarkAsDone(context.Background(), uuid.New())
		require.Error(t, err)
	})
		t.Run("todo doesn't eixst", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().MarkAsDone(gomock.Any(), gomock.Any()).Return(customerr.ERR_TODO_NOT_EXIST)

		s := NewService(repos)
		err := s.MarkAsDone(context.Background(), uuid.New())
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
}

func TestService_DeleteTodoByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().DeleteTodo(gomock.Any(), gomock.Any()).Return(nil)

		s := NewService(repos)
		err := s.DeleteTodoByID(context.Background(), uuid.New())
		require.NoError(t, err)
	})

	t.Run("returns error", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().DeleteTodo(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

		s := NewService(repos)
		err := s.DeleteTodoByID(context.Background(), uuid.New())
		require.Error(t, err)
	})
	t.Run("todo doesn't eixst", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().DeleteTodo(gomock.Any(), gomock.Any()).Return(customerr.ERR_TODO_NOT_EXIST)

		s := NewService(repos)
		err := s.DeleteTodoByID(context.Background(), uuid.New())
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
}

func TestService_GetAllTodos(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		want := []entity.Todo{}
		
		todo1 := entity.NewTodo(testDeadline, "dad's birthday", testUserID)
		want = append(want, todo1)
		todo2 := entity.NewTodo(testDeadline, "bake a cake", testUserID)
		want = append(want, todo2)
		todo3 := entity.NewTodo(testDeadline, "get up early", testUserID)
		want = append(want, todo3)

		repos := newRepos(t)
		repos.EXPECT().GetAllTodos(gomock.Any(), gomock.Any()).Return(want, nil)

		s := NewService(repos)
		got, err := s.GetAllTodos(context.Background(), testUserID)
		
		require.NoError(t, err)
		require.EqualValues(t, want, got)
	})

	t.Run("returns error", func(t *testing.T) {
		repos := newRepos(t)
		repos.EXPECT().GetAllTodos(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))

		s := NewService(repos)
		got, err := s.GetAllTodos(context.Background(), testUserID)
		require.Error(t, err)

		require.Equal(t, []entity.Todo{}, got)
	})
}

func TestService_UpdateTodosBody(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().UpdateTodosBody(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		s := NewService(repo)

		err := s.UpdateTodosBody(context.Background(), uuid.New(), "body")
		require.NoError(t, err)
	})
	t.Run("todo doesn't exist", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().UpdateTodosBody(gomock.Any(), gomock.Any(), gomock.Any()).Return(customerr.ERR_TODO_NOT_EXIST)
		s := NewService(repo)

		err := s.UpdateTodosBody(context.Background(), uuid.New(), "body")
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
	t.Run("internal error", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().UpdateTodosBody(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		s := NewService(repo)

		err := s.UpdateTodosBody(context.Background(), uuid.New(), "body")
		require.Error(t, err)
	})
}

func TestService_UpdateTodosDeadline(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().UpdateTodosDeadline(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		s := NewService(repo)

		err := s.UpdateTodosDeadline(context.Background(), uuid.New(), testDeadline)
		require.NoError(t, err)
	})
	t.Run("todo doesn't exist", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().UpdateTodosDeadline(gomock.Any(), gomock.Any(), gomock.Any()).Return(customerr.ERR_TODO_NOT_EXIST)
		s := NewService(repo)

		err := s.UpdateTodosDeadline(context.Background(), uuid.New(), testDeadline)
		require.ErrorIs(t, err, customerr.ERR_TODO_NOT_EXIST)
	})
	t.Run("internal error", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().UpdateTodosDeadline(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		s := NewService(repo)

		err := s.UpdateTodosDeadline(context.Background(), uuid.New(), testDeadline)
		require.Error(t, err)
	})
}

func TestService_DeleteDoneTodos(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().DeleteDoneTodos(gomock.Any(), gomock.Any()).Return(nil)
		s := NewService(repo)

		err := s.DeleteDoneTodos(context.Background(), uuid.New())
		require.NoError(t, err)
	})
	
	t.Run("internal error", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().DeleteDoneTodos(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		s := NewService(repo)

		err := s.DeleteDoneTodos(context.Background(), uuid.New())
		require.Error(t, err)
	})
}

func TestService_DeletePassedDeadline(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().DeletePassedDeadline(gomock.Any(), gomock.Any()).Return(nil)
		s := NewService(repo)

		err := s.DeletePassedDeadline(context.Background(), uuid.New())
		require.NoError(t, err)
	})
	
	t.Run("internal error", func(t *testing.T) {
		repo := newRepos(t)
		repo.EXPECT().DeletePassedDeadline(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		s := NewService(repo)

		err := s.DeletePassedDeadline(context.Background(), uuid.New())
		require.Error(t, err)
	})
}