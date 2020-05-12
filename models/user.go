package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
