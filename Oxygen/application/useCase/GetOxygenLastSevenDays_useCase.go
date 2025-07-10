package useCase

import "github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"

type GetOxygenLastSevenDaysUc struct {
	db domain.IOxygenRepository
}

func NewGetOxygenLastSevenDaysUc(db domain.IOxygenRepository) *GetOxygenLastSevenDaysUc {
	return &GetOxygenLastSevenDaysUc{db: db}
}
func (useCase *GetOxygenLastSevenDaysUc) Execute(userId int) (map[string][]domain.GetOxygenModel, error) {
	return useCase.db.GetLastSevenDays(userId)
}
