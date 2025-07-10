package main

import (
  heartRateDependencies "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/dependencies"
  heartRatesRoutes "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/routes"
	oxygenDependencies "github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/dependencies"
	oxygenRoutes "github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/routes"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
	temperatureDependencies "github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/dependencies"
	temperatureRoutes "github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	oxygenDependencies.Init()
  temperatureDependencies.Init()
  heartRateDependencies.Init()
	

	// Create a new Gin router
	r := gin.Default()
	helpers.InitCORS(r)

	temperatureRoutes.Routes(r)
  oxygenRoutes.Routes(r)
  heartRatesRoutes.Routes(r)

	r.Run(":8080") // Start the server on port 8080

}

