package service


import (
	"github.com/bchanona/vitals_warmheart_backend/Temperature/domain"
)

type DailyAverageService struct {
	Data map[string][]domain.GetTemperatureModel
}

func NewDailyAverageService(data map[string][]domain.GetTemperatureModel) *DailyAverageService {
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
