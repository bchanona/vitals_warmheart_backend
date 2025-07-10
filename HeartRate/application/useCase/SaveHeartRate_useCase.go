package useCase

import "github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"

type SaveHeartRate struct {
	HeartRateRepository domain.IHeartRateRepository
}

func NewSaveHeartRate(heartRateRepository domain.IHeartRateRepository) *SaveHeartRate {
	return &SaveHeartRate{
		HeartRateRepository: heartRateRepository,
	}
}
func (s *SaveHeartRate) Execute(heartRate domain.SaveHeartRateModel) error {
	return s.HeartRateRepository.Save(heartRate)
}
