package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wtran29/event-booking/models"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
