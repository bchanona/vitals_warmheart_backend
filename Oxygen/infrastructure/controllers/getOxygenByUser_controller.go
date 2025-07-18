package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetOxygenByUserController struct {
	useCase *useCase.GetOxygenByUserUc
}

func NewGetOxygenByUserController(useCase *useCase.GetOxygenByUserUc) *GetOxygenByUserController {
	return &GetOxygenByUserController{useCase: useCase}
}
func (controller *GetOxygenByUserController) Execute(ctx *gin.Context) {
	userIdstr, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
	}

	userId, ok := userIdstr.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	oxygens, err := controller.useCase.Execute(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Oxygens": oxygens})
}
