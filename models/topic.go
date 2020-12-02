package models

import "gorm.io/gorm"

// Topic model
type Topic struct {
	gorm.Model
	Name string `json:name`
	// Question []Question
}
