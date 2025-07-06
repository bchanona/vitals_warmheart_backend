package useCase

import "github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"

type GetOxygenByDateUc struct {
	db domain.IOxygenRepository
}

func NewGetOxygenByDateUc(db domain.IOxygenRepository) *GetOxygenByDateUc {
	return &GetOxygenByDateUc{db: db}
}
func (useCase *GetOxygenByDateUc) Execute(userId int, date string) ([]domain.GetOxygenModel, error) {
	return useCase.db.GetByDate(userId, date)
}