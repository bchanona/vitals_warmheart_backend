package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/bchanona/vitals_warmheart_backend/Temperature/domain"
	"github.com/bchanona/vitals_warmheart_backend/Temperature/queries"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) Save(temperature domain.SaveTemperatureModel) error {
	_, err := sql.db.Exec(queries.SaveTemperatureQuery, temperature.User_id, temperature.Measurement, temperature.Device_id)
	return err
}

func (sql *MySQL) GetAll() ([]domain.GetTemperatureModel, error) {
	var temperatures []domain.GetTemperatureModel

	rows, err := sql.db.Query(queries.GetAllTemperatures)

	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var temp domain.GetTemperatureModel

		err := rows.Scan(
			&temp.Measurement,
			&temp.Date,
			&temp.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		temperatures = append(temperatures, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return temperatures, nil

}

func (sql *MySQL) GetByDate(user_id int, date string) ([]domain.GetTemperatureModel, error) {
	var temperaturesByDate []domain.GetTemperatureModel

	rows, err := sql.db.Query(queries.GetByDate, user_id, date)

	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var temperatureByD domain.GetTemperatureModel

		err := rows.Scan(
			&temperatureByD.Measurement,
			&temperatureByD.Date,
			&temperatureByD.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		temperaturesByDate = append(temperaturesByDate, temperatureByD)

	}
	return temperaturesByDate, nil

}

func (sql *MySQL) GetByUser(user_id int) ([]domain.GetTemperatureModel, error) {
	var temperaturesByUser []domain.GetTemperatureModel

	rows, err := sql.db.Query(queries.GetByUser, user_id)

	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var temperatureByUser domain.GetTemperatureModel
		err := rows.Scan(
			&temperatureByUser.Measurement,
			&temperatureByUser.Date,
			&temperatureByUser.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		temperaturesByUser = append(temperaturesByUser, temperatureByUser)
	}
	return temperaturesByUser, nil

}

func (sql *MySQL) GetForSupervisor(user_id int) ([]domain.GetTemperatureByUserModel, error) {
	var temps []domain.GetTemperatureByUserModel

	rows, err := sql.db.Query(queries.GetForSupervisor, user_id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tempForSu domain.GetTemperatureByUserModel
		err := rows.Scan(
			&tempForSu.Measurement,
			&tempForSu.Date,
			&tempForSu.Time,
			&tempForSu.Name_user,
			&tempForSu.Surname_user,
			&tempForSu.Email_user,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		temps = append(temps, tempForSu)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return temps, nil
}
func (sql *MySQL) GetLast7Days(user_id int)(map[string][]domain.GetTemperatureModel, error){
	result := make(map[string][]domain.GetTemperatureModel)

	rows, err := sql.db.Query(queries.GetLast7Days, user_id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var temp domain.GetTemperatureModel
		err := rows.Scan(&temp.Measurement, &temp.Date, &temp.Time)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		
		date := temp.Date 
		result[date] = append(result[date], temp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return result, nil
}
