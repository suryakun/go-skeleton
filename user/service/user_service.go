package service

import (
	"context"

	"github.com/suryakun/skeleton-go/models"
	"github.com/suryakun/skeleton-go/user"
)

type userService struct {
	userRepo user.Repository
}

// NewUserService ...
func NewUserService(ur user.Repository) user.Service {
	return &userService{
		userRepo: ur,
	}
}

func (s *userService) Store(c context.Context, user models.User) error {
	err := s.userRepo.CreateUser(c, user)
	return err
}

func (s *userService) All(c context.Context, limit, offset int) ([]models.User, error) {
	users, err := s.userRepo.GetUsers(c, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) Get(c context.Context, id int64) (*models.User, error) {
	user, err := s.userRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
