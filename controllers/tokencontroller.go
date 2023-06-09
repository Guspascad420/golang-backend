package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test/auth"
	"test/database"
	"test/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meta": &models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Db.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, "Invalid credentials"}})
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "Success"}, "data": models.LoginResponse{tokenString}})
}

func ExtractToken(c *gin.Context) (string, error) {
	token := c.Query("token")
	if token != "" {
		return token, nil
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("Unauthorized Error: Access Denied")
}

func ExtractEmail(c *gin.Context) (string, error) {
	signedToken, err := ExtractToken(c)
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&auth.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("F)J@NcQfTjWnZr4u7x!A%D*G-KaPdSgV"), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*auth.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}
	return claims.Email, nil
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
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{false, "success"}, "data": user})
}
