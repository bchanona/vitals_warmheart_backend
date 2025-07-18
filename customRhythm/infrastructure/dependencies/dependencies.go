package dependencies

import (
	"log"

	"github.com/bchanona/vitals_warmheart_backend/customRhythm/application/usecase"
	"github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure"
	"github.com/bchanona/vitals_warmheart_backend/customRhythm/infrastructure/controllers"
	"github.com/bchanona/vitals_warmheart_backend/helpers"
)

var (
	mySQL infrastructure.MySQL
)

func Init() {
	db, err := helpers.ConnectMySQL()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL = *infrastructure.NewMySQL(db)

}

func GetRegisterCustomRhythm() *controllers.SaveCustomController{
	useCase := usecase.NewSaveCustomRhytmUc(&mySQL)
	return controllers.NewSaveCustomController(useCase)
}
func GetUpdateCustomRhythm() *controllers.UpdateCustomRhytmController{
	useCase := usecase.NewUpdateCustomRhytmUc(&mySQL)
	return controllers.NewUpdateCustomRhytmController(useCase)
}