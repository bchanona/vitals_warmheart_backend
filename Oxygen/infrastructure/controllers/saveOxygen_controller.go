package controllers

import (
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/useCase"
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"
	"github.com/gin-gonic/gin"
)

type SaveOxygenController struct {
	useCase *useCase.SaveOxygenUc
}

func NewSaveOxygenController(useCase *useCase.SaveOxygenUc) *SaveOxygenController {
	return &SaveOxygenController{useCase: useCase}
}

func (controller *SaveOxygenController) Execute(ctx *gin.Context) {
	var oxygen domain.SaveOxygenModel

	if err := ctx.ShouldBindJSON(&oxygen); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := controller.useCase.Execute(oxygen); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save oxygen data"})
	}
	ctx.JSON(200, gin.H{"message": "Oxygen data saved successfully"})

}
