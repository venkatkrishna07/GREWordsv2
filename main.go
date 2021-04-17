package main

import (
	"GREWordsv2/controllers"

	"GREWordsv2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/word", controllers.GetWord)
	r.GET("/learnt", controllers.Learntwords)
	r.GET("/check/:word", controllers.CheckWord)
	r.GET("/getW/:word", controllers.GetMeaning)
	r.GET("/check", controllers.Testfunc)
	r.POST("/addWord", controllers.AddWord)

	r.Run()
}
