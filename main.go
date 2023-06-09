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
		c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "success"}})
	})
	r.POST("/user/register", controllers.RegisterUser)
	r.POST("/user/login", controllers.GenerateToken)
	r.GET("/:building/rooms", controllers.GetAllRooms)
	r.POST("/user/bookings/create", controllers.CreateBooking)
	r.GET("/bookings", controllers.GetBookingsByDate)
	r.GET("/user/profile", controllers.GetUserProfile)
	r.GET("/user/bookings", controllers.GetUserBookings)
	r.PUT("/user/bookings/update/:id", controllers.EditBooking)
	r.DELETE("/user/bookings/delete/:id", controllers.DeleteBooking)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
