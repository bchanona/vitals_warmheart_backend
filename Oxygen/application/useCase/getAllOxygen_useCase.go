package useCase

import "github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"

type GetAllOxygenUc struct {
	db domain.IOxygenRepository
}

func NewGetAllOxygenUc(db domain.IOxygenRepository) *GetAllOxygenUc {
	return &GetAllOxygenUc{db: db}
}
func (useCase *GetAllOxygenUc) Execute() ([]domain.GetOxygenModel, error) {
	return useCase.db.GetAll()
}