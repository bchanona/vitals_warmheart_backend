package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/Temperature/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetTempByUserController struct {
	useCase *useCase.GetTempByUserUc
}

func NewGetTempByUserController(useCase *useCase.GetTempByUserUc) *GetTempByUserController {
	return &GetTempByUserController{useCase: useCase}
}

func (controller *GetTempByUserController) Execute(ctx *gin.Context) {
	userIdstr, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id de usuario no proporcionado"})
	}

	userId, ok := userIdstr.(int)

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
