package domain

// ITemperatureRepository defines the interface for temperature data operations.
type ITemperatureRepository interface {
	//Save Temperature
	Save(temperature SaveTemperatureModel) error
	//Get all recorded temperatures
	GetAll() ([]GetTemperatureModel, error)
	//Get temperature by date
	GetByDate(user_id int, date string) ([]GetTemperatureModel, error)
	//Get temperature by user
	GetByUser(user_id int) ([]GetTemperatureModel, error)
	//Get temperature for supervisor
	GetForSupervisor(user_id int) ([]GetTemperatureByUserModel, error)
	//Get temperature for the last 7 days
	GetLast7Days(user_id int)(map[string][]GetTemperatureModel, error)
}
