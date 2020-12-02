package controllers

import (
	"quiz-api/db"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

// GetAll : get all topic
func GetAllTopic(c *gin.Context) {
	var topics []models.Topic
	db := db.Dbconn()
	result := db.Find(&topics)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"data": topics,
		})
	}
}
