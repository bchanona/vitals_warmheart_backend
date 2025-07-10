package controllers

import (
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/useCase"
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"
	"github.com/gin-gonic/gin"
)

type SaveHeartRateController struct {
	useCase *useCase.SaveHeartRate
}

func NewSaveHeartRateController(useCase *useCase.SaveHeartRate) *SaveHeartRateController {
	return &SaveHeartRateController{
		useCase: useCase,
	}
}
func (controller *SaveHeartRateController) Execute(ctx *gin.Context) {
	var heartRateModel domain.SaveHeartRateModel
	if err := ctx.ShouldBindJSON(&heartRateModel); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := controller.useCase.Execute(heartRateModel)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save heart rate"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Heart rate saved successfully"})
}
