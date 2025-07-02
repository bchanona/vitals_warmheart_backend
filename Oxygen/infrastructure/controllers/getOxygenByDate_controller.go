package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetOxygenByDateController struct {
	useCase *useCase.GetOxygenByDateUc
}

func NewGetOxygenByDateController(useCase *useCase.GetOxygenByDateUc) *GetOxygenByDateController {
	return &GetOxygenByDateController{useCase: useCase}
}
func (controller *GetOxygenByDateController) Execute(ctx *gin.Context) {
	user_id_str, exist := ctx.Get("user_id")

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id de usuario no proporcionado"})
	}

	user_id, ok := user_id_str.(int)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}
	date := ctx.Param("date")
	oxygenByDate, err := controller.useCase.Execute(user_id, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": oxygenByDate})
}
