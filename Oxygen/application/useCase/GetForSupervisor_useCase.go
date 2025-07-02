package useCase

import "github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"

type GetForSupervisorUc struct {
	db domain.IOxygenRepository
}

func NewGetForSupervisorUc(db domain.IOxygenRepository) *GetForSupervisorUc {
	return &GetForSupervisorUc{db: db}
}
func (useCase *GetForSupervisorUc) Execute(userId int) ([]domain.GetOxygenByUserModel, error) {
	return useCase.db.GetForSupervisor(userId)
}
