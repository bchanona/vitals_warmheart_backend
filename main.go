package main

import (
	temperatureDependencies "github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/dependencies"
	temperatureRoutes "github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/routes"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	temperatureDependencies.Init()
	

	// Create a new Gin router
	r := gin.Default()
	helpers.InitCORS(r)
	temperatureRoutes.Routes(r)

	r.Run(":8080") // Start the server on port 8080
}
