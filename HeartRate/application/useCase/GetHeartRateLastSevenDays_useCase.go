package useCase

import "github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"

type GetHeartRateLastSevenDays struct {
	HeartRateRepository domain.IHeartRateRepository
}

func NewGetHeartRateLastSevenDays(heartRateRepository domain.IHeartRateRepository) *GetHeartRateLastSevenDays {
	return &GetHeartRateLastSevenDays{
		HeartRateRepository: heartRateRepository,
	}
}
func (useCase *GetHeartRateLastSevenDays) Execute(userId int) (map[string][]domain.GetHeartRateModel, error) {
	return useCase.HeartRateRepository.GetLast7Days(userId)
}