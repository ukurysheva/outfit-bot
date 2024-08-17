package service

import (
	"testing"

	"outfitbot/internal/model"
)

func TestService_PrecipitationRecommendation(t *testing.T) {
	tests := []struct {
		name    string
		weather *model.Weather
		want    string
	}{
		{
			name:    "precipitation more than 30%",
			weather: &model.Weather{Daily: model.WeatherDaily{PrecipitationProbability: []int{90}}},
			want:    "Не забудьте зонт!",
		},
		{
			name:    "precipitation less than 30%",
			weather: &model.Weather{Daily: model.WeatherDaily{PrecipitationProbability: []int{20}}},
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}

			if got := s.PrecipitationRecommendation(tt.weather); got != tt.want {
				t.Errorf("PrecipitationRecommendation() = %v, want %v", got, tt.want)
			}
		})
	}
}
