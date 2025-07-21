package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/queries"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) Save(heartRate domain.SaveHeartRateModel) error {
	_, err := sql.db.Exec(queries.SaveHeartRateQuery, heartRate.User_id, heartRate.Measurement, heartRate.Device_id)
	return err
}
func (sql *MySQL) GetAll() ([]domain.GetHeartRateModel, error) {
	var heartRates []domain.GetHeartRateModel

	rows, err := sql.db.Query(queries.GetAllHeartRates)

	if err != nil {
		return nil, fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var heartRate domain.GetHeartRateModel

		err := rows.Scan(
			&heartRate.Measurement,
			&heartRate.Date,
			&heartRate.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		heartRates = append(heartRates, heartRate)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while traversing the row")
	}
	return heartRates, nil
}
func (sql *MySQL) GetByDate(userId int, date string) ([]domain.GetHeartRateModel, error) {
	var heartRatesByDate []domain.GetHeartRateModel
	rows, err := sql.db.Query(queries.GetByDate, userId, date)
	if err != nil {
		return nil, fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var heartRateByDate domain.GetHeartRateModel
		err := rows.Scan(
			&heartRateByDate.Measurement,
			&heartRateByDate.Date,
			&heartRateByDate.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		heartRatesByDate = append(heartRatesByDate, heartRateByDate)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while traversing the row")
	}
	return heartRatesByDate, nil
}
func (sql *MySQL) GetByUser(userId int) ([]domain.GetHeartRateModel, error) {
	var heartRatesByUser []domain.GetHeartRateModel
	rows, err := sql.db.Query(queries.GetByUser, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var heartRateByUser domain.GetHeartRateModel
		err := rows.Scan(
			&heartRateByUser.Measurement,
			&heartRateByUser.Date,
			&heartRateByUser.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		heartRatesByUser = append(heartRatesByUser, heartRateByUser)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while traversing the row")
	}
	return heartRatesByUser, nil
}
func (sql *MySQL) GetForSupervisor(userId int) ([]domain.GetHeartRateByUserModel, error) {
	var heartRAtesForSupervisor []domain.GetHeartRateByUserModel
	rows, err := sql.db.Query(queries.GetForSupervisor, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var heartRateForSupervisor domain.GetHeartRateByUserModel
		err := rows.Scan(
			&heartRateForSupervisor.Measurement,
			&heartRateForSupervisor.Date,
			&heartRateForSupervisor.Time,
			&heartRateForSupervisor.Name_user,
			&heartRateForSupervisor.Surname_user,
			&heartRateForSupervisor.Email_user,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		heartRAtesForSupervisor = append(heartRAtesForSupervisor, heartRateForSupervisor)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while traversing the row")
	}
	return heartRAtesForSupervisor, nil
}
func (sql *MySQL) GetLast7Days(userId int) (map[string][]domain.GetHeartRateModel, error) {
	heartRatesByLast7Days := make(map[string][]domain.GetHeartRateModel)

	rows, err := sql.db.Query(queries.GetLast7Days, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var heartRate domain.GetHeartRateModel
		err := rows.Scan(
			&heartRate.Measurement,
			&heartRate.Date,
			&heartRate.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		heartRatesByLast7Days[heartRate.Date] = append(heartRatesByLast7Days[heartRate.Date], heartRate)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while traversing the row")
	}
	return heartRatesByLast7Days, nil
}
