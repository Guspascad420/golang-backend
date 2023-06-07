package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/database"
	"test/models"
)

func GetAllRooms(c *gin.Context) {
	building := c.Param("building")
	var rooms []models.Room
	record := database.Db.Where("building = ?", building).Find(&rooms)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": &models.Meta{false, record.Error.Error()}})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": &models.Meta{true, "Success"}, "data": rooms})
}
