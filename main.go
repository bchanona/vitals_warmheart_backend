package main

import "github.com/bchanona/vitals_warmheart_backend/helpers"

func main() {
	//Verificamos la conexión a la base de datos
	db,_ := helpers.ConnectMySQL()
	if db == nil {
		panic("Error connecting to the database")
	}else {
		println("Successful connection to the database")
	}

}