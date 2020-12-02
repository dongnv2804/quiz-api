package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	UserName string
	FullName string
	Password string
}
