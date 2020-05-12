package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/suryakun/skeleton-go/models"
)

// Service ...
type Service struct {
	mock.Mock
}

// Store ...
func (s *Service) Store(c context.Context, user models.User) error {
	args := s.Called(c, user)
	return args.Error(0)
}

// All ...
func (s *Service) All(c context.Context, limit, offset int) ([]models.User, error) {
	args := s.Called(c, limit, offset)
	return args.Get(0).([]models.User), args.Error(1)
}

// Get ...
func (s *Service) Get(c context.Context, id int64) (*models.User, error) {
	args := s.Called(c, id)
	return args.Get(0).(*models.User), args.Error(1)
}
