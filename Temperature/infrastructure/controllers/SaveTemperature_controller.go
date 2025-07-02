package controllers

import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/application/useCase"
	"github.com/bchanona/vitals_warmheart_backend/Temperature/domain"
	"github.com/gin-gonic/gin"
)

type SaveTemperatureController struct {
	useCase *useCase.SaveTemperatureUseCase
}

func NewSaveTemperatureController(useCase *useCase.SaveTemperatureUseCase) *SaveTemperatureController {
	return &SaveTemperatureController{useCase: useCase}
}
func (controller *SaveTemperatureController) Execute(ctx *gin.Context){
	var temperature domain.SaveTemperatureModel

	if err := ctx.ShouldBindJSON(&temperature); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	err := controller.useCase.Execute(temperature)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save temperature"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Temperature saved successfully"})

}