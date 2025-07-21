package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/service"
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetHeartRateLastSevenDaysController struct {
	useCase *useCase.GetHeartRateLastSevenDays
}

func NewGetHeartRateLastSevenDaysController(useCase *useCase.GetHeartRateLastSevenDays) *GetHeartRateLastSevenDaysController {
	return &GetHeartRateLastSevenDaysController{
		useCase: useCase,
	}
}
func (controller *GetHeartRateLastSevenDaysController) Execute(ctx *gin.Context) {
	user_id_str, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found in context"})
	}

	user_id, ok := user_id_str.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}
	heartRates, err := controller.useCase.Execute(user_id)
	service := service.NewDailyAverageService(heartRates)
	averages := service.Calculate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve heart rates"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Averages: ": averages})

}
