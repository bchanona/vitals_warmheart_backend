package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetHeartRateByUserController struct {
	useCase *useCase.GetHeartRateByUser
}
func NewGetHeartRateByUserController(useCase *useCase.GetHeartRateByUser) *GetHeartRateByUserController {
	return &GetHeartRateByUserController{
		useCase: useCase,
	}
}
func (controller *GetHeartRateByUserController) Execute(ctx *gin.Context){
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve heart rates"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"heartRates": heartRates})
}