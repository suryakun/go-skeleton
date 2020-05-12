package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/suryakun/skeleton-go/models"
	"github.com/suryakun/skeleton-go/user"
)

type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) user.Repository {
	db.AutoMigrate(&models.User{})
	return &userRepository{
		DB: db,
	}
}

func (p *userRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	user := new(models.User)
	err := p.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userRepository) GetUsers(ctx context.Context, limit, offset int) ([]models.User, error) {
	var users []models.User
	err := p.DB.Find(&users).Limit(limit).Offset(offset).Error
	return users, err
}

func (p *userRepository) CreateUser(ctx context.Context, user models.User) error {
	err := p.DB.Create(&user).Error
	return err
}
