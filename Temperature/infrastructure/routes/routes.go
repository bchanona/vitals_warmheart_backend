package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/temperature")

	saveTemperature := dependencies.GetSaveTemperatureController().Execute

	routes.POST("/", saveTemperature)
}
