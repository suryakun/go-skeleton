package user

import (
	"context"

	"github.com/suryakun/skeleton-go/models"
)

// Service ...
type Service interface {
	Store(c context.Context, user models.User) error
	All(c context.Context, limit, offset int) ([]models.User, error)
	Get(c context.Context, id int64) (*models.User, error)
}
