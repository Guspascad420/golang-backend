package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/database"
	"test/models"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meta": models.Meta{Message: err.Error()}})
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
