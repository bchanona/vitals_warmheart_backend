package useCase

import "github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"

type GetHeartRateByDate struct {
	HeartRateRepository domain.IHeartRateRepository
}

func NewGetHeartRateByDate(heartRateRepository domain.IHeartRateRepository) *GetHeartRateByDate {
	return &GetHeartRateByDate{
		HeartRateRepository: heartRateRepository,
	}
}
func (useCase *GetHeartRateByDate) Execute(userId int, date string) ([]domain.GetHeartRateModel, error) {
	return useCase.HeartRateRepository.GetByDate(userId, date)
}