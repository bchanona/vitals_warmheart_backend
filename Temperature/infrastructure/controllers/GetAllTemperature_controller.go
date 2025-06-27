package controllers

import (
	"net/http"

	"github.com/bchanona/vitals_warmheart_backend/Temperature/application/useCase"
	"github.com/gin-gonic/gin"
)

type GetAllTemperaturesController struct {
	useCase *useCase.GetAllTemperaturesUc
}

func NewGetAllTemperaturesController(useCase *useCase.GetAllTemperaturesUc) *GetAllTemperaturesController{
	return &GetAllTemperaturesController{useCase: useCase}
}

func (controller *GetAllTemperaturesController) Execute(ctx *gin.Context){

	temperatures, err := controller.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"Temperatures": temperatures})
}

