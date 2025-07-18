package main

import (
	heartRateDependencies "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/dependencies"
	customRhythmDependencies "github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure/dependencies"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
	"github.com/gin-gonic/gin"
	heartRatesRoutes "github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/routes"
	customRhythmRoutes "github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure/routes" 
)

func main() {
	
	heartRateDependencies.Init()
	customRhythmDependencies.Init()

	r:= gin.Default()
	helpers.InitCORS(r)

	heartRatesRoutes.Routes(r)
	customRhythmRoutes.Routes(r)

	r.Run(":8080") 


}