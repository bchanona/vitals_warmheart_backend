package useCase

import "github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"

type GetAllHeartRate struct {
	HeartRateRepository domain.IHeartRateRepository
}

func NewGetAllHeartRate(heartRateRepository domain.IHeartRateRepository) *GetAllHeartRate {
	return &GetAllHeartRate{
		HeartRateRepository: heartRateRepository,
	}
}
func (useCase *GetAllHeartRate) Execute() ([]domain.GetHeartRateModel, error) {
	return useCase.HeartRateRepository.GetAll()
}
