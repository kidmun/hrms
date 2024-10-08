package controllers

import (
	"hrms/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase models.UserUsecase
}


func (uc *UserController) CreateUser(ctx *gin.Context){
	var userInput models.UserRequest
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	

}