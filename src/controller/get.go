package controller

import (
	"aqua-ratsnake/src/database"
	"aqua-ratsnake/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var Users []model.User
	if err := database.Database.Where(("id = ?"), c.Param("id")).First(&Users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Users})
}