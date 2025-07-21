package main

import (
	customRhythmDependencies "github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure/dependencies"
  heartRateDependencies "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/dependencies"
	oxygenDependencies "github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/dependencies"
  temperatureDependencies "github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/dependencies"
	temperatureRoutes "github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/routes"
  heartRatesRoutes "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/routes"
	oxygenRoutes "github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/routes"
  customRhythmRoutes "github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure/routes"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
	"github.com/gin-gonic/gin"
)


func main() {
	  oxygenDependencies.Init()
  	temperatureDependencies.Init()
  	heartRateDependencies.Init()
  	customRhythmDependencies.Init()
	


	// Create a new Gin router
	r := gin.Default()
	helpers.InitCORS(r)


	customRhythmRoutes.Routes(r)
	temperatureRoutes.Routes(r)
  oxygenRoutes.Routes(r)
	heartRatesRoutes.Routes(r)

	r.Run(":8080") // Start the server on port 8080

}

