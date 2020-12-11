package models

import "gorm.io/gorm"

// Answer model
type Answer struct {
	gorm.Model
	Content    string
	IsCorrect  bool
	QuestionID uint
}
