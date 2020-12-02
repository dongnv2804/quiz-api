package controllers

import (
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
	// var answers []models.Answer
	db := db.Dbconn()

	result := db.Where("topicid = ?", c.Param("topicid")).Find(&questions)
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

type Dataquestion struct {
	Id     uint `json:id`
	Answer bool `json:answer`
}

type DataPost struct {
	Dataquestion []Dataquestion `json:dataquestion`
}

// CaculatorScore : caculate score question
func CaculatorScore(c *gin.Context) {
	db := db.Dbconn()
	var question models.Question
	var data DataPost
	var score int32 = 0
	if err := c.BindJSON(&data); err == nil {
		for _, v := range data.Dataquestion {
			result := db.Where("ID = ? and answer = ?", v.Id, v.Answer).First(&question)
			if result.Error != nil {
				panic(result.Error)
			} else {
				score += question.Score
			}
		}
		defer c.JSON(200, gin.H{
			"score": score,
		})
	} else {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

}
