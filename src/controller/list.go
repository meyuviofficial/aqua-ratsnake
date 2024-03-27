package controller

import (
	"aqua-ratsnake/src/database"
	"aqua-ratsnake/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
var Users []model.User
database.Database.Find(&Users)

c.JSON(http.StatusOK, gin.H{"data": Users})
}