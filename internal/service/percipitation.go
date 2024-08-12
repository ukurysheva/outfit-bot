package service

import (
	"fmt"

	"outfitbot/internal/model"
)

var precipitationTmpl = "🌧 Вероятность осадков %d%%."

func (s *Service) PrecipitationInfo(weather *model.Weather) string {
	return fmt.Sprintf(precipitationTmpl, weather.Daily.PrecipitationProbability[0])
}

func (s *Service) PrecipitationRecommendation(weather *model.Weather) string {
	if weather.Daily.PrecipitationProbability[0] > 50 {
		return "Не забудьте зонт!"
	}

	return ""
}
