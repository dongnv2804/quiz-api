package models

import "gorm.io/gorm"

// Topic model
type Topic struct {
	gorm.Model
	Name string
	// Questions []Question
}
