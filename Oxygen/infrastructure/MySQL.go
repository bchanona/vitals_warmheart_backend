package infrastructure

import (
	"database/sql"

	"github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"
	"github.com/bchanona/vitals_warmheart_backend/queries"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) Save(oxygen domain.SaveOxygenModel) error {
	_, err := sql.db.Exec(queries.SaveOxygen, oxygen.User_id, oxygen.Measurement, oxygen.Device_id)
	return err
}
func (sql *MySQL) GetAll() ([]domain.GetOxygenModel, error) {
	rows, err := sql.db.Query(queries.GetAllOxygen)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var oxygenData []domain.GetOxygenModel
	for rows.Next() {
		var oxygen domain.GetOxygenModel
		if err := rows.Scan(&oxygen.Measurement, &oxygen.Date, &oxygen.Time); err != nil {
			return nil, err
		}
		oxygenData = append(oxygenData, oxygen)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return oxygenData, nil
}
func (sql *MySQL) GetByDate(userId int, date string) ([]domain.GetOxygenModel, error) {
	rows, err := sql.db.Query(queries.GetOxygenByDate, userId, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var oxygenByDate []domain.GetOxygenModel
	for rows.Next() {
		var oxygenBD domain.GetOxygenModel
		if err := rows.Scan(&oxygenBD.Measurement, &oxygenBD.Date, &oxygenBD.Time); err != nil {
			return nil, err
		}
		oxygenByDate = append(oxygenByDate, oxygenBD)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return oxygenByDate, nil
}
