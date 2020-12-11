package controllers

import (
	"fmt"
	"quiz-api/db"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

// AddQuestion : add question to db
func AddQuestion(c *gin.Context) {
	db := db.Dbconn()
	var question models.Question
	db.AutoMigrate(&question)
	if err := c.ShouldBindJSON(&question); err == nil {
		result := db.Create(&question)
		if result.RowsAffected > 0 {
			c.JSON(201, gin.H{
				"message": "create question successfull",
			})
		} else {
			c.JSON(500, gin.H{
				"err": result.Error,
			})
		}
	} else {
		c.JSON(500, gin.H{
			"err": err,
		})
	}
}

// GetQuestionByTopicID : get all question by topicID
func GetQuestionByTopicID(c *gin.Context) {
	var questions []models.Question
	db := db.Dbconn()
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.Answer{})
	result := db.Preload("Answers").Where("topic_id = ?", c.Param("topicid")).Find(&questions)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"data": questions,
		})
	}
}

// Dataquestion is struct
type Dataquestion struct {
	QuestionID uint `json:questionId`
	AnswerID   uint `json:answerId`
}

// DataPost is struct
type DataPost struct {
	Dataquestion []Dataquestion `json:dataquestion`
}

// CaculatorScore : caculate score question
func CaculatorScore(c *gin.Context) {
	db := db.Dbconn()
	var data DataPost
	var score, countCorrect int32 = 0, 0
	if err := c.BindJSON(&data); err == nil {
		for _, v := range data.Dataquestion {
			var question models.Question
			result := db.Preload("Answers", "id = ?", v.AnswerID).Where("id = ?", v.QuestionID).Find(&question)
			if result.Error != nil {
				fmt.Println(result.Error)
			}
			if question.Answers[0].IsCorrect != false {
				score += question.Score
				countCorrect++
			}
		}
		c.JSON(200, gin.H{
			"score":        score,
			"countCorrect": countCorrect,
		})

	} else {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

}
