package dependencies

import (
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/application/useCase"
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure"
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/infrastructure/controllers"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
)

var (
	mySQL infrastructure.MySQL
)

func Init(){
	//Verificamos la conexión a la base de datos
	db,_ := helpers.ConnectMySQL()
	if db == nil {
		panic("Error connecting to the database")
	}else {
		println("Successful connection to the database")
	}

	mySQL = *infrastructure.NewMySQL(db)
}

func GetSaveHeartRateController() *controllers.SaveHeartRateController{
	useCase := useCase.NewSaveHeartRate(&mySQL)
	return controllers.NewSaveHeartRateController(useCase)
}

func GetGetAllHeartRateController() *controllers.GetAllHeartRateController{
	useCase := useCase.NewGetAllHeartRate(&mySQL)
	return controllers.NewGetAllHeartRateController(useCase)
}
func GetHeartRateBydateController() *controllers.GetHeartRateByDateController{
	useCase := useCase.NewGetHeartRateByDate(&mySQL)
	return controllers.NewGetHeartRateByDateController(useCase)
}
func GetHeartRateByUserController() *controllers.GetHeartRateByUserController{
	useCase := useCase.NewGetHeartRateByUser(&mySQL)
	return controllers.NewGetHeartRateByUserController(useCase)
}
func GetForSupervisorController() *controllers.GetForSupervisorController{
	useCase := useCase.NewGetHeartRateForSupervisor(&mySQL)
	return controllers.NewGetForSupervisorController(useCase)
}
func GetHeartRateLastSevenDaysController() *controllers.GetHeartRateLastSevenDaysController{
	useCase := useCase.NewGetHeartRateLastSevenDays(&mySQL)
	return controllers.NewGetHeartRateLastSevenDaysController(useCase)
}