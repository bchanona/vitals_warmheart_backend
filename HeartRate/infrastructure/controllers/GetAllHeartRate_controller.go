package controllers

import (
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetAllHeartRateController struct {
	useCase *useCase.GetAllHeartRate
}

func NewGetAllHeartRateController(useCase *useCase.GetAllHeartRate) *GetAllHeartRateController {
	return &GetAllHeartRateController{
		useCase: useCase,
	}
}
func (controller *GetAllHeartRateController) Execute(ctx *gin.Context) {
	heartRates, err := controller.useCase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve heart rates"})
		return
	}
	ctx.JSON(200, gin.H{"heartRates": heartRates})

}
