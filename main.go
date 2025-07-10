package main

import (
	heartRateDependencies "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/dependencies"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
	"github.com/gin-gonic/gin"
	heartRatesRoutes "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/routes"
)

func main() {
	
	heartRateDependencies.Init()

	r:= gin.Default()
	helpers.InitCORS(r)

	heartRatesRoutes.Routes(r)

	r.Run(":8080") 


}