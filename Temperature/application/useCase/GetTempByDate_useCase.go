package useCase

import "github.com/bchanona/vitals_warmheart_backend/Temperature/domain"

type GetTempByDateUc struct {
	db domain.ITemperatureRepository
}

func NewGetTempByDateUc (db domain.ITemperatureRepository) *GetTempByDateUc{
	return &GetTempByDateUc{db: db}
}

func (uc *GetTempByDateUc) Execute(user_id int, date string)([]domain.GetTemperatureModel, error){
	return uc.db.GetByDate(user_id, date)
}