package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/database"
	"test/models"
	"time"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	record := database.Db.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func GetUserProfile(c *gin.Context) {
	var user models.User
	email, err := ExtractEmail(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	record := database.Db.Where("email = ?", email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"meta": models.Meta{true, "success"},
		"data": models.UserProfileResponse{user.ID, user.Name, user.Username, user.Email}})
}

func GetUserBookings(c *gin.Context) {
	userId := c.Query("userId")
	var bookings []models.Booking

	record := database.Db.Where("user_id = ?", userId).Where("date > ?", time.Now()).Preload("Room").Find(&bookings)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "success"}, "data": bookings})
}
