package server

import (
	"context"
	"log"
	"testing"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	customErr "github.com/NeverlandMJ/ToDo/user-service/pkg/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestServer_CreateUser(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("user exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		err = s.CreateUser(context.Background(), user)
		assert.ErrorIs(t, err, customErr.ERR_USER_EXIST)
	})

	t.Run("success", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)
	})
}

func TestServer_GetUser(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("user doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))
		
		user := entity.NewUser("sunbula", "123", "+123456789")
		
		u, err := s.GetUser(context.Background(), user.UserName, user.Password)
		require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
		assert.Equal(t, entity.User{}, u)
	})

	t.Run("success", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		u, err := s.GetUser(context.Background(), user.UserName, user.Password)
		require.NoError(t, err)
		assert.Equal(t, user, u)
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
		"file://./../migrations",
	)

	require.NoError(t, err)

	return serv
}

func cleanUpFn(s *Server) func() {
	return func() {
		if err := s.deleteUsers(); err != nil {
			log.Println("CLEANUP OF DB FAILED!")
		}
	}
}
