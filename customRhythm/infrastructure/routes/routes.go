package routes

import (
	"github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure/dependencies"
	"github.com/bchanona/vitals_warmheart_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/customRhythm")
	saveCustom := dependencies.GetRegisterCustomRhythm().Execute
	updateCustom := dependencies.GetUpdateCustomRhythm().Execute

	routes.POST("/",middlewares.AuthMiddleware(),saveCustom)
	routes.PUT("/update",middlewares.AuthMiddleware(),updateCustom)
}