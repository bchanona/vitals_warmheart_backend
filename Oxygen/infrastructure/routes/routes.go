package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/dependencies"
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
	routes.GET("/:date", getOxygenByDateController)
	routes.GET("/user", getOxygenByUserController)
	routes.GET("/supervisor", getForSupervisor)
	routes.GET("/lastSevenDays", getOxygenLastSevenDaysController)


}