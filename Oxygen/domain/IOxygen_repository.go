package domain

type IOxygenRepository interface{
	Save(oxygen SaveOxygenModel) error

	GetAll()([]GetOxygenModel, error)

	GetByDate(userId int, date string)([]GetOxygenModel, error)

	GetByUser(userId int)([]GetOxygenModel, error)

	GetForSupervisor(userId int)([]GetOxygenByUserModel, error)

	GetLastSevenDays(userId int)(map[string][]GetOxygenModel, error)
}