package service

import "outfitbot/internal/model"

const (
	uvIndexMediumLevel = 3
	uvIndexHighLevel   = 5
	uvIndexExtraLevel  = 8
)

func (s *Service) UvIndexRecommendation(weather *model.Weather) string {
	if len(weather.Daily.UvIndexMax) == 0 {
		return ""
	}

	uvIndex := weather.Daily.UvIndexMax[0]
	rec := ""

	if uvIndex >= uvIndexExtraLevel {
		rec = "Обязательно нанесите SPF 30+, наденьте головной убор и избегайте прямых лучей."
	} else if uvIndex >= uvIndexHighLevel {
		rec = "Обязательно нанесите SPF 15+ и наденьте головной убор."
	} else if uvIndex >= uvIndexMediumLevel {
		rec = "Рекомендуем нанести SPF 15+."
	}

	return rec
}
