package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/service"
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetOxygenLastSevenDaysController struct {
	useCase *useCase.GetOxygenLastSevenDaysUc
}

func NewGetOxygenLastSevenDaysController(useCase *useCase.GetOxygenLastSevenDaysUc) *GetOxygenLastSevenDaysController {
	return &GetOxygenLastSevenDaysController{useCase: useCase}
}
func (controller *GetOxygenLastSevenDaysController) Execute(ctx *gin.Context) {
	userIdstr, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id de usuario no proporcionado"})
	}

	userId, ok := userIdstr.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	oxygens, err := controller.useCase.Execute(userId)
	service := service.NewDailyAverageService(oxygens)
	averages := service.Calculate()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Averages: ": averages})
}