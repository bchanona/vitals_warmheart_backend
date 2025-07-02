package dependencies

import (
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/application/useCase"
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure"
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/infrastructure/controllers"
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

func GetSaveOxygenController() *controllers.SaveOxygenController {
	useCase := useCase.NewSaveOxygenUc(&mySQL)
	return controllers.NewSaveOxygenController(useCase)
}
func GetGetAllOxygenController() *controllers.GetAllOxygenController {
	useCase := useCase.NewGetAllOxygenUc(&mySQL)
	return controllers.NewGetAllOxygenController(useCase)
}
func GetGetOxygenByDateController() *controllers.GetOxygenByDateController {
	useCase := useCase.NewGetOxygenByDateUc(&mySQL)
	return controllers.NewGetOxygenByDateController(useCase)
}

