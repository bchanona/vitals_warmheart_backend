package useCase

import "github.com/bchanona/vitals_warmheart_backend/Oxygen/domain"

type SaveOxygenUc struct {
	db domain.IOxygenRepository
}

func NewSaveOxygenUc (db domain.IOxygenRepository) *SaveOxygenUc{
	return &SaveOxygenUc{db: db}
}

func (useCase *SaveOxygenUc)Execute(oxygen domain.SaveOxygenModel) error {
	return useCase.db.Save(oxygen)
}