package service

import (
	"context"

	"github.com/suryakun/skeleton-go/models"
	"github.com/suryakun/skeleton-go/user"
)

type userService struct {
	UserRepo user.Repository
}

// NewUserService ...
func NewUserService(ur user.Repository) user.Service {
	return &userService{
		UserRepo: ur,
	}
}

func (s *userService) Store(c context.Context, user models.User) error {
	err := s.UserRepo.CreateUser(c, user)
	return err
}

func (s *userService) All(c context.Context, limit, offset int) ([]models.User, error) {
	users, err := s.UserRepo.GetUsers(c, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) Get(c context.Context, id int64) (*models.User, error) {
	user, err := s.UserRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
