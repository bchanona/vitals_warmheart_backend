package useCase

import "github.com/bchanona/vitals_warmheart_backend/Temperature/domain"

type GetAllTemperaturesUc struct {
	db domain.ITemperatureRepository
}

func NewGetAllTemperaturesUc(db domain.ITemperatureRepository) *GetAllTemperaturesUc {
	return &GetAllTemperaturesUc{db: db}
}

func (uc *GetAllTemperaturesUc) Execute() ([]domain.GetTemperatureModel, error) {
	return uc.db.GetAll()
}
