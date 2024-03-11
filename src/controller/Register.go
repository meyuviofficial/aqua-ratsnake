package controller

import (
	"aqua-ratsnake/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input model.User

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Password:       input.Password,
		RegistrationId: input.RegistrationId,
		DateOfBirth:    input.DateOfBirth,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while saving the user", "message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedUser})
}
