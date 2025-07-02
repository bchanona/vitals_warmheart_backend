package useCase

import "github.com/bchanona/vitals_warmheart_backend/Temperature/domain"

//Define use case to save temperature

type SaveTemperatureUseCase struct {
	db domain.ITemperatureRepository
}

func NewSaveTemperatureUseCase(db domain.ITemperatureRepository) *SaveTemperatureUseCase {
	return &SaveTemperatureUseCase{
		db: db,
	}
}
// Execute saves the temperature data using the repository.
func (useCase *SaveTemperatureUseCase) Execute(temperature domain.SaveTemperatureModel) error {
	return useCase.db.Save(temperature)
}