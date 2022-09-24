package server

import (
	"context"
	"log"
	"testing"

	"github.com/NeverlandMJ/ToDo/user-service/config"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/customErr"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/google/uuid"
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

		user := entity.NewUser("sunbula", "123", "+998123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		u, err := s.GetUser(context.Background(), user.UserName, user.Password)
		user.Password = u.Password
		require.NoError(t, err)
		assert.Equal(t, user, u)
	})
}

func TestServer_ChangePassword(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("user doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		err := s.ChangePassword(context.Background(), uuid.New(), "old", "new")
		require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
	})

	t.Run("success", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+998123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		err = s.ChangePassword(context.Background(), user.ID, user.Password, "456")
		require.NoError(t, err)

		_, err = s.GetUser(context.Background(), user.UserName, "456")
		require.NoError(t, err)
	})

	t.Run("incorrect old password", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+998123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		err = s.ChangePassword(context.Background(), user.ID, "old", "456")
		require.ErrorIs(t, err, customErr.ERR_INCORRECT_PASSWORD)

		_, err = s.GetUser(context.Background(), user.UserName, "456")
		require.ErrorIs(t, err, customErr.ERR_INCORRECT_PASSWORD)
	})
}

func TestServer_ChangeUserName(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	t.Run("user doesn't exist", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		err := s.ChangeUserName(context.Background(), uuid.New(), "123")
		require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
	})

	t.Run("success", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+998123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		err = s.ChangeUserName(context.Background(), user.ID, "neverland")
		require.NoError(t, err)

		got, err := s.GetUser(context.Background(), "neverland", "123")
		require.NoError(t, err)
		require.Equal(t, got.UserName, "neverland")
	})

	t.Run("user name is already taken", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user1 := entity.NewUser("sunbula", "123", "+998123456789")
		err := s.CreateUser(context.Background(), user1)
		require.NoError(t, err)

		user2 := entity.NewUser("neverland", "123", "+99888563479")
		err = s.CreateUser(context.Background(), user2)
		require.NoError(t, err)

		err = s.ChangeUserName(context.Background(), user1.ID, "neverland")
		require.Error(t, err, err.Error())
	})
}

func TestServer_DeleteAccount(t *testing.T) {
	s := newServer(t)
	t.Cleanup(cleanUpFn(s))

	// t.Run("user doens't eixts", func(t *testing.T) {
	// 	t.Cleanup(cleanUpFn(s))

	// 	err := s.DeleteAccount(context.Background(), uuid.New(), "123", "user")
	// 	require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
	// })

	t.Run("success", func(t *testing.T) {
		t.Cleanup(cleanUpFn(s))

		user := entity.NewUser("sunbula", "123", "+998123456789")
		err := s.CreateUser(context.Background(), user)
		require.NoError(t, err)

		err = s.DeleteAccount(context.Background(), user.ID, user.Password, user.UserName)
		require.NoError(t, err)

		got, err := s.GetUser(context.Background(), user.UserName, user.Password)
		require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
		require.EqualValues(t, entity.User{}, got)
	})

	// t.Run("user's password is incorrect", func(t *testing.T) {
	// 	t.Cleanup(cleanUpFn(s))

	// 	user := entity.NewUser("sunbula", "123", "+998123456789")
	// 	err := s.CreateUser(context.Background(), user)
	// 	require.NoError(t, err)

	// 	err = s.DeleteAccount(context.Background(), user.ID, "635", user.UserName)
	// 	require.ErrorIs(t, err, customErr.ERR_INCORRECT_PASSWORD)
	// })

	// t.Run("user name is incorrect", func(t *testing.T) {
	// 	t.Cleanup(cleanUpFn(s))

	// 	user := entity.NewUser("sunbula", "123", "+998123456789")
	// 	err := s.CreateUser(context.Background(), user)
	// 	require.NoError(t, err)

	// 	err = s.DeleteAccount(context.Background(), user.ID, user.Password, "user")
	// 	require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
	// })

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
			PostgresMigrationsPath: "file://./../database/migrations",
		},
		
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
