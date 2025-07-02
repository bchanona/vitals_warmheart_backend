package controllers

import (
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetAllOxygenController struct {
	useCase *useCase.GetAllOxygenUc
}

func NewGetAllOxygenController(useCase *useCase.GetAllOxygenUc) *GetAllOxygenController {
	return &GetAllOxygenController{useCase: useCase}
}
func (controller *GetAllOxygenController) Execute(ctx *gin.Context) {
	oxygenData, err := controller.useCase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve oxygen data"})
		return
	}
	ctx.JSON(200, gin.H{"data": oxygenData})
}
