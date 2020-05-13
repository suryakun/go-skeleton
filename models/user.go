package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
