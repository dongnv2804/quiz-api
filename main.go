package main

import (
	"quiz-api/controllers"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	r := setupRoute()
	r.Run(":8080")
}

func setupRoute() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "wellcome api of Dongkisot",
			})
		})
		api.GET("/topic", controllers.GetAllTopic)
		api.GET("/question/:topicid", controllers.GetQuestionByTopicID)
		api.POST("/question/create", controllers.AddQuestion)
		api.POST("/getscore", controllers.CaculatorScore)
	}
	return r
}
