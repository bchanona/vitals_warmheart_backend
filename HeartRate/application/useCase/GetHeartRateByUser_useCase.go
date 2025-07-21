package useCase

import "github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"

type GetHeartRateByUser struct {
	HeartRateRepository domain.IHeartRateRepository
}

func NewGetHeartRateByUser(heartRateRepository domain.IHeartRateRepository) *GetHeartRateByUser {
	return &GetHeartRateByUser{
		HeartRateRepository: heartRateRepository,
	}
}
func (useCase *GetHeartRateByUser) Execute(userId int) ([]domain.GetHeartRateModel, error) {
	return useCase.HeartRateRepository.GetByUser(userId)
}