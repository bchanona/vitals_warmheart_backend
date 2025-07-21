package service


import (
	"github.com/bchanona/vitals_warmheart_backend/HeartRate/domain"
)

type DailyAverageService struct {
	Data map[string][]domain.GetHeartRateModel
}

func NewDailyAverageService(data map[string][]domain.GetHeartRateModel) *DailyAverageService {
	return &DailyAverageService{
		Data: data,
	}
}

func (s *DailyAverageService) Calculate() map[string]float64 {
	averages := make(map[string]float64)

	for date, records := range s.Data {
		var total float64
		for _, record := range records {
			total += record.Measurement
		}
		if len(records) > 0 {
			averages[date] = total / float64(len(records))
		}
	}

	return averages
}