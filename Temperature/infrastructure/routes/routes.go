package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/dependencies"
	"github.com/bchanona/vitals_warmheart_backend/middlewares"
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
	routes.GET("/:date",middlewares.AuthMiddleware(),getByDate)
	routes.GET("/user",middlewares.AuthMiddleware(),getByUser)
	routes.GET("/supervisor",middlewares.AuthSupervisorMiddleware(), getForSupervisor)
	routes.GET("/lastSevenDays",middlewares.AuthMiddleware(), getTemperaturesLast7Days)
}
