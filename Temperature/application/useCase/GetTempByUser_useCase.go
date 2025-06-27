package useCase

import "github.com/bchanona/vitals_warmheart_backend/Temperature/domain"

type GetTempByUserUc struct {
	db domain.ITemperatureRepository
}

func NewGetTempByUserUc (db domain.ITemperatureRepository) *GetTempByUserUc{
	return &GetTempByUserUc{db: db}
}

func (uc *GetTempByUserUc) Execute(user_id int)([]domain.GetTemperatureModel, error){
	return uc.db.GetByUser(user_id)
}