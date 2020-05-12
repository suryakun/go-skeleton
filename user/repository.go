package user

import (
	"context"

	"github.com/suryakun/skeleton-go/models"
)

// Repository User
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetUsers(ctx context.Context, limit, offset int) ([]models.User, error)
	CreateUser(ctx context.Context, user models.User) error
}
