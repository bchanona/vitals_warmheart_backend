package infrastructure

import (
	"database/sql"

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