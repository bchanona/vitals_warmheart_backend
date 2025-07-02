package main

import (
	oxygenDependencies "github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/dependencies"
	oxygenRoutes "github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	oxygenDependencies.Init()

	r:= gin.Default()
	oxygenRoutes.Routes(r)

	r.Run(":8080")


}
