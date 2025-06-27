package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/Temperature/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetForSupervisorController struct {
	useCase *useCase.GetForSupervisorUc
}

func NewGetForSupervisorController(useCase *useCase.GetForSupervisorUc) *GetForSupervisorController {
	return &GetForSupervisorController{useCase: useCase}
}

func (controller *GetForSupervisorController) Execute(ctx *gin.Context) {
	userIdStr, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id de usuario no proporcionado"})
	}

	userId, ok := userIdStr.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}
	temperatures, err := controller.useCase.Execute(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Temperatures": temperatures})
}
