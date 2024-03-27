package controller

import (
	"aqua-ratsnake/src/database"
	"aqua-ratsnake/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var Users []model.User
	if err := database.Database.Where(("id = ?"), c.Param("id")).Find(&Users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
  	var input model.UserUpdateInput
  	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  database.Database.Model(&Users).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": Users})
	
}