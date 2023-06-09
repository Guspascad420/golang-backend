package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test/database"
	"test/models"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meta": models.Meta{Message: err.Error() + "sknmdjfenwe"}})
		c.Abort()
		return
	}
	record := database.Db.Create(&booking)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{Message: record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"meta": models.Meta{true, "Success"}, "data": nil})
}

func GetBookingsByDate(c *gin.Context) {
	date := strings.Replace(c.Query("date"), "-", "/", 2)
	var bookings []models.Booking
	record := database.Db.Where("date = ?", date).Preload("Room").Find(&bookings)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "Success"}, "data": bookings})
}
