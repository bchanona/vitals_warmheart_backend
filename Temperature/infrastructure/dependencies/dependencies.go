package dependencies

import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/infrastructure"
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
