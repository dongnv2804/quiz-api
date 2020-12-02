package models

import "gorm.io/gorm"

// Question model
type Question struct {
	gorm.Model
	Content string `json:content`
	Score   int32  `json:score`
	Answer  bool   `json:answer`
	TopicId uint   `json:toppicid gorm:"column:topic_id"`
}
