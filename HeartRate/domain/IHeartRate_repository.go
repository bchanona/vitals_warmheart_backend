package domain

type IHeartRateRepository interface {
	//Save Temperature
	Save(heartRate SaveHeartRateModel) error
	GetAll() ([]GetHeartRateModel, error)
	GetByDate(userId int, date string) ([]GetHeartRateModel, error)
	GetByUser(userId int) ([]GetHeartRateModel, error)
	GetForSupervisor(userId int) ([]GetHeartRateByUserModel, error)
	GetLast7Days(userId int) (map[string][]GetHeartRateModel, error)
}
