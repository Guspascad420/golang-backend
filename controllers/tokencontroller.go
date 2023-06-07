package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/auth"
	"test/database"
	"test/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"meta": &models.Meta{false, err.Error()}})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Db.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"meta": &models.Meta{false, "Invalid credentials"}})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, err.Error()}})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"meta": &models.Meta{true, "Success"}, "token": tokenString})
}
