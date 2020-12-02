package controllers

import (
	"quiz-api/db"
	"quiz-api/models"

	"github.com/gin-gonic/gin"
)

// Login : user login
func Login(c *gin.Context) {
	var user models.User
	db := db.Dbconn()
	if err := c.ShouldBindJSON(&user); err == nil {
		result := db.Where("UserName = ? and Password = ?", user.UserName, user.Password).First(&user)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": result.Error,
			})
		}
		if user.ID != 0 {
			c.JSON(200, gin.H{
				"message": "login ok",
			})
		}

	} else {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

}
