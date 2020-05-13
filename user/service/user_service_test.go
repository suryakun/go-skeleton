package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suryakun/skeleton-go/models"
	"github.com/suryakun/skeleton-go/user/mocks"
	"github.com/suryakun/skeleton-go/user/service"
)

func TestStore(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	mockUser := &models.User{
		Name:  "surya",
		Phone: "1234123",
		Email: "testing@test.cxv",
	}

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mockUserRepo.On("GetByID", ctx, int64(1)).Return(mockUser, nil)
		userService := service.NewUserService(mockUserRepo)
		user, err := userService.Get(ctx, int64(1))
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ctx := context.Background()
		mockUserRepo.On("GetByID", ctx, int64(2)).Return(nil, models.ErrNotFound)
		userService := service.NewUserService(mockUserRepo)
		user, err := userService.Get(ctx, int64(2))
		assert.Nil(t, user)
		assert.NotNil(t, err)
	})
}

func TestAll(t *testing.T) {
	var users []models.User
	mockUserRepo := new(mocks.Repository)
	mockUser := &models.User{
		Name:  "surya",
		Phone: "1234123",
		Email: "testing@test.cxv",
	}
	users = append(users, *mockUser)

	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()

		mockUserRepo.On("GetUsers", ctx, 1, 2).Return(users, nil)
		userService := service.NewUserService(mockUserRepo)
		user, err := userService.All(ctx, 1, 2)
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		ctx := context.Background()
		mockUserRepo.On("GetUsers", ctx, 2, 2).Return(nil, models.ErrNotFound)
		userService := service.NewUserService(mockUserRepo)
		users, err := userService.All(ctx, 2, 2)
		assert.Nil(t, users)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	mockUser := &models.User{
		Name:  "surya",
		Phone: "1234123",
		Email: "testing@test.cxv",
	}

	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		mockUserRepo.On("GetByID", ctx, int64(1)).Return(mockUser, nil)
		userService := service.NewUserService(mockUserRepo)
		user, err := userService.Get(ctx, 1)
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		ctx := context.Background()
		mockUserRepo.On("GetByID", ctx, int64(2)).Return(nil, models.ErrNotFound)
		userService := service.NewUserService(mockUserRepo)
		user, err := userService.Get(ctx, int64(2))
		assert.NotNil(t, err)
		assert.Nil(t, user)
	})
}
