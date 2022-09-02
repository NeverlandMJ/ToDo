package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/NeverlandMJ/ToDo/user-service/mocks"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/customErr"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestService_CreateUser(t *testing.T) {
	t.Run("should return error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

		s := service.NewService(mockRepo)
		user := entity.NewUser("sunbula", "123", "+123456789")
		got := s.CreateUser(context.Background(), user)

		require.Error(t, got)
	})

	t.Run("should pass", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

		s := service.NewService(mockRepo)
		user := entity.NewUser("sunbula", "123", "+123456789")
		got := s.CreateUser(context.Background(), user)

		require.NoError(t, got)
	})
}

func TestService_GetUser(t *testing.T) {
	t.Run("should return error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().GetUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(entity.User{}, fmt.Errorf("error"))

		s := service.NewService(mockRepo)
		got, err := s.GetUser(context.Background(), "sunbula", "123")

		require.Error(t, err)
		require.EqualValues(t, entity.User{}, got)
	})

	t.Run("should pass", func(t *testing.T) {
		want := entity.User{
			ID:        uuid.New(),
			UserName:  "sunbula",
			Password:  "213",
			Phone:     "+123456789",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsBlocked: true,
		}
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().GetUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(want, nil)

		s := service.NewService(mockRepo)
		got, err := s.GetUser(context.Background(), "sunbula", "213")

		require.NoError(t, err)
		require.EqualValues(t, want, got)
	})
}

func TestService_ChangePassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().ChangePassword(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		s := service.NewService(mockRepo)
		err := s.Repo.ChangePassword(context.Background(), uuid.New(), "old", "new")
		require.NoError(t, err)
	})
	t.Run("incorrect old password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().ChangePassword(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(customErr.ERR_INCORRECT_PASSWORD)

		s := service.NewService(mockRepo)
		err := s.Repo.ChangePassword(context.Background(), uuid.New(), "old", "new")
		require.ErrorIs(t, err, customErr.ERR_INCORRECT_PASSWORD)
	})

	t.Run("user doesn't exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().ChangePassword(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(customErr.ERR_USER_NOT_EXIST)

		s := service.NewService(mockRepo)
		err := s.Repo.ChangePassword(context.Background(), uuid.New(), "old", "new")
		require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
	})
}

func TestService_ChangeUserName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().ChangeUserName(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		s := service.NewService(mockRepo)
		err := s.Repo.ChangeUserName(context.Background(), uuid.New(), "new")
		require.NoError(t, err)
	})

	t.Run("user doesn't exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().ChangeUserName(gomock.Any(), gomock.Any(), gomock.Any()).Return(customErr.ERR_USER_NOT_EXIST)

		s := service.NewService(mockRepo)
		err := s.Repo.ChangeUserName(context.Background(), uuid.New(), "new")
		require.ErrorIs(t, err, customErr.ERR_USER_NOT_EXIST)
	})
}
