package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/temperature")

	saveTemperature := dependencies.GetSaveTemperatureController().Execute
	getTemperatures := dependencies.GetAllTemperaturesController().Execute
	getByDate := dependencies.GetTempByDate().Execute
	getByUser := dependencies.GetTempByUser().Execute
	getForSupervisor := dependencies.GetForSupervisor().Execute
	getTemperaturesLast7Days := dependencies.GetTempLast7DaysController().Execute
	

	routes.POST("/", saveTemperature)
	routes.GET("/", getTemperatures)
	routes.GET("/:date",getByDate)
	routes.GET("/user",getByUser)
	routes.GET("/supervisor", getForSupervisor)
	routes.GET("/lastSevenDays", getTemperaturesLast7Days)
}
