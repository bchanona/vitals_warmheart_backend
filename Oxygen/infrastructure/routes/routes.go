package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/dependencies"
	"github.com/bchanona/vitals_warmheart_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/oxygen")
	
	saveOxygenController := dependencies.GetSaveOxygenController().Execute
	getAllOxygenController := dependencies.GetGetAllOxygenController().Execute
	getOxygenByDateController := dependencies.GetGetOxygenByDateController().Execute
	getOxygenByUserController := dependencies.GetOxygenByUserController().Execute
	getForSupervisor := dependencies.GetForSupervisorController().Execute
	getOxygenLastSevenDaysController := dependencies.GetGetOxygenLastSevenDaysController().Execute

	routes.POST("/",saveOxygenController)
	routes.GET("/",getAllOxygenController)
	routes.GET("/:date",middlewares.AuthMiddleware(), getOxygenByDateController)
	routes.GET("/user",middlewares.AuthMiddleware(), getOxygenByUserController)
	routes.GET("/supervisor",middlewares.AuthSupervisorMiddleware(), getForSupervisor)
	routes.GET("/lastSevenDays",middlewares.AuthMiddleware(), getOxygenLastSevenDaysController)


}