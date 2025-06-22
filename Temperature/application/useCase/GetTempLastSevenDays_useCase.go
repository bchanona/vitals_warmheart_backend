package useCase

import "github.com/bchanona/vitals_warmheart_backend/Temperature/domain"

type GetTempLastSevenDaysUc struct {
	db domain.ITemperatureRepository
}

func NewGetTempLastSevenDaysUc(db domain.ITemperatureRepository) *GetTempLastSevenDaysUc{
	return &GetTempLastSevenDaysUc{db: db}
}

func (useCase *GetTempLastSevenDaysUc) Execute(user_id int)([]domain.GetTemperatureModel, error){
	return useCase.db.GetLast7Days(user_id)
}