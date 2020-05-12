package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/suryakun/skeleton-go/models"
)

// Repository ...
type Repository struct {
	mock.Mock
}

// GetByID ...
func (_m *Repository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	args := _m.Called(ctx, id)
	if id == 2 {
		return nil, args.Error(1)
	}
	return args.Get(0).((*models.User)), args.Error(1)
}

// GetUsers ...
func (_m *Repository) GetUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	args := _m.Called(ctx, limit, offset)
	if limit == 2 {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

// CreateUser ...
func (_m *Repository) CreateUser(ctx context.Context, user models.User) error {
	args := _m.Called(ctx, user)
	return args.Error(0)
}
