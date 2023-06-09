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
		context.JSON(http.StatusUnauthorized, gin.H{"meta": models.Meta{false, "Invalid credentials"}})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, err.Error()}})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"meta": models.Meta{true, "Success"}, "data": models.LoginResponse{tokenString}})
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractEmail(c *gin.Context) (string, error) {
	signedToken := ExtractToken(c)
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

func GetEmail(c *gin.Context) {
	email, err := ExtractEmail(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": models.Meta{false, err.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": models.Meta{false, "success"}, "data": email})
}
