package useCase

import "github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"

type GetHeartRateForSupervisor struct {
	HeartRateRepository domain.IHeartRateRepository
}

func NewGetHeartRateForSupervisor(heartRateRepository domain.IHeartRateRepository) *GetHeartRateForSupervisor {
	return &GetHeartRateForSupervisor{
		HeartRateRepository: heartRateRepository,
	}
}
func (useCase *GetHeartRateForSupervisor) Execute(userId int) ([]domain.GetHeartRateByUserModel, error) {
	return useCase.HeartRateRepository.GetForSupervisor(userId)
}