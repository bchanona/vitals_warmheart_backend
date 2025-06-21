package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/bchanona/vitals_warmheart_backend/Temperature/domain"
	"github.com/bchanona/vitals_warmheart_backend/queries"
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
			&temp.Temperature_id,
			&temp.User_id,
			&temp.Measurement,
			&temp.Date,
			&temp.Time,
			&temp.Device_id,	
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
