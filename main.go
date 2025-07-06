package main

import (
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
	

	// Create a new Gin router
	r := gin.Default()
	helpers.InitCORS(r)

	temperatureRoutes.Routes(r)
    oxygenRoutes.Routes(r)

	r.Run(":8080") // Start the server on port 8080

}
