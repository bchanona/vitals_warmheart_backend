package dependencies

import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/application/useCase"
	"github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure"
	"github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure/controllers"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
)

var (
	mySQL infrastructure.MySQL
)

func Init(){
	db,_ := helpers.ConnectMySQL()
	if db == nil {
		panic("Error connecting to the database")
	}else {
		println("Successful connection to the database")
	}
	mySQL = *infrastructure.NewMySQL(db)
}



func GetSaveTemperatureController() *controllers.SaveTemperatureController{
	useCase := useCase.NewSaveTemperatureUseCase(&mySQL)
	return controllers.NewSaveTemperatureController(useCase)
}
func GetAllTemperaturesController() *controllers.GetAllTemperaturesController{
	useCase := useCase.NewGetAllTemperaturesUc(&mySQL)
	return controllers.NewGetAllTemperaturesController(useCase)
}
func GetTempByDate() *controllers.GetTempByDateController {
	useCase := useCase.NewGetTempByDateUc(&mySQL)
	return  controllers.NewGetTempByDateController(useCase)
}
func GetTempByUser() *controllers.GetTempByUserController{
	useCase := useCase.NewGetTempByUserUc(&mySQL)
	return  controllers.NewGetTempByUserController(useCase)
}
func GetForSupervisor() *controllers.GetForSupervisorController{
	useCase := useCase.NewGetForSupervisorUc(&mySQL)
	return controllers.NewGetForSupervisorController(useCase)
}
func GetTempLast7DaysController() *controllers.GetTempLastSevenDaysController{
	useCase := useCase.NewGetTempLastSevenDaysUc(&mySQL)
	return controllers.NewGetTempLastSevenDaysController(useCase)
}