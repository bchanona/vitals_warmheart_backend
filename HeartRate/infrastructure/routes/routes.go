package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/dependencies"
	"github.com/bchanona/vitals_warmheart_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/heartRate")
	saveHeartRate := dependencies.GetSaveHeartRateController().Execute
	getAllHeartRates := dependencies.GetGetAllHeartRateController().Execute
	getByDate := dependencies.GetHeartRateBydateController().Execute
	getByUser := dependencies.GetHeartRateByUserController().Execute
	getForSupervisor := dependencies.GetForSupervisorController().Execute
	getLast7Days := dependencies.GetHeartRateLastSevenDaysController().Execute

	routes.POST("/",saveHeartRate)
	routes.GET("/", getAllHeartRates)
	routes.GET("/:date",middlewares.AuthMiddleware(), getByDate)
	routes.GET("/user",middlewares.AuthMiddleware(), getByUser)
	routes.GET("/supervisor",middlewares.AuthSupervisorMiddleware(), getForSupervisor)
	routes.GET("/last7days",middlewares.AuthMiddleware(), getLast7Days)

	
}