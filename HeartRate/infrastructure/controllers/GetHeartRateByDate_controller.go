package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetHeartRateByDateController struct {
	useCase *useCase.GetHeartRateByDate
}

func NewGetHeartRateByDateController(useCase *useCase.GetHeartRateByDate) *GetHeartRateByDateController {
	return &GetHeartRateByDateController{
		useCase: useCase,
	}
}
func (controller *GetHeartRateByDateController) Execute(ctx *gin.Context) {
	user_id_str, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found in context"})
	}

	user_id, ok := user_id_str.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}
	date := ctx.Param("date")
	heartRates, err := controller.useCase.Execute(user_id, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve heart rates"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"heartRates": heartRates})
}
