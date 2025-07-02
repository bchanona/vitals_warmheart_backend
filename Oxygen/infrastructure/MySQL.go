package infrastructure

import (
	"database/sql"

	"github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"
	"github.com/bchanona/vitals_warmheart_backend/Oxygen/queries"
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
func (sql *MySQL) GetByUser(userId int) ([]domain.GetOxygenModel, error) {
	var oxygenByUser []domain.GetOxygenModel
	rows, err := sql.db.Query(queries.GetByUser, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var oxygenBU domain.GetOxygenModel
		if err := rows.Scan(&oxygenBU.Measurement, &oxygenBU.Date, &oxygenBU.Time); err != nil {
			return nil, err
		}
		oxygenByUser = append(oxygenByUser, oxygenBU)
	}
	return oxygenByUser, nil
}
func (sql *MySQL) GetForSupervisor(userId int) ([]domain.GetOxygenByUserModel, error) {
	rows, err := sql.db.Query(queries.GetForSupervisor, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var oxygenForSupervisor []domain.GetOxygenByUserModel
	for rows.Next() {
		var oxygenFS domain.GetOxygenByUserModel
		if err := rows.Scan(&oxygenFS.Measurement, &oxygenFS.Date, &oxygenFS.Time, &oxygenFS.Name_user, &oxygenFS.Surname_user, &oxygenFS.Email_user); err != nil {
			return nil, err
		}
		oxygenForSupervisor = append(oxygenForSupervisor, oxygenFS)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return oxygenForSupervisor, nil
}
func (sql *MySQL) GetLastSevenDays(userId int) (map[string][]domain.GetOxygenModel, error) {
	reult := make(map[string][]domain.GetOxygenModel)
	rows, err := sql.db.Query(queries.GetLastSevenDays, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var oxygenLD domain.GetOxygenModel
		if err := rows.Scan(&oxygenLD.Measurement, &oxygenLD.Date, &oxygenLD.Time); err != nil {
			return nil, err
		}
		reult[oxygenLD.Date] = append(reult[oxygenLD.Date], oxygenLD)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reult, nil

}
