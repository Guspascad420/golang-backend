package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/controllers"
	"test/database"
	"test/models"
)

func main() {
	r := gin.Default()

	database.Connect()
	database.Migrate()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/meta", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"meta": &models.Meta{true, "success"}})
	})
	r.POST("/register", controllers.RegisterUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
