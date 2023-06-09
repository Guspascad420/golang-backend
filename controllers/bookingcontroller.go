package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

func EditBooking(c *gin.Context) {
	var booking models.Booking
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meta": models.Meta{Message: err.Error()}})
		c.Abort()
		return
	}

	update := database.Db.Model(models.Booking{ID: uint(id)}).Updates(booking)
	if update.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, update.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "successfully updated"}, "data": nil})
}

func DeleteBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleteRec := database.Db.Delete(&models.Booking{}, id)
	if deleteRec.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, deleteRec.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "successfully deleted"}, "data": nil})
}
