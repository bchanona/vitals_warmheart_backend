package useCase

import "github.com/bchanona/vitals_warmheart_backend/Temperature/domain"

type GetForSupervisorUc struct {
	db domain.ITemperatureRepository
}

func NewGetForSupervisorUc(db domain.ITemperatureRepository) *GetForSupervisorUc {
	return &GetForSupervisorUc{db: db}
}

func (uc *GetForSupervisorUc) Execute(user_id int) ([]domain.GetTemperatureByUserModel, error) {
	return uc.db.GetForSupervisor(user_id)
}