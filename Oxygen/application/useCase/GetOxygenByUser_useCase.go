package useCase

import "github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"

type GetOxygenByUserUc struct {
	db domain.IOxygenRepository
}

func NewGetOxygenByUserUc(db domain.IOxygenRepository) *GetOxygenByUserUc {
	return &GetOxygenByUserUc{db: db}
}
func (useCase *GetOxygenByUserUc) Execute(userId int) ([]domain.GetOxygenModel, error) {
	return useCase.db.GetByUser(userId)
}
