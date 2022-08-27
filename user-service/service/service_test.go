package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/NeverlandMJ/ToDo/user-service/mocks"
	"github.com/NeverlandMJ/ToDo/user-service/pkg/entity"
	"github.com/NeverlandMJ/ToDo/user-service/service"
	"github.com/golang/mock/gomock"
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
			ID: "some_id",
			UserName: "sunbula",
			Password: "213",
			Phone: "+123456789",
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
