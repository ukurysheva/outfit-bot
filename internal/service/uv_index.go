package service

import "outfitbot/internal/model"

const (
	uvIndexMediumLevel = 3
	uvIndexExtraLevel  = 8
)

func (s *Service) UvIndexRecommendation(weather *model.Weather) string {
	if len(weather.Daily.UvIndexMax) == 0 {
		return ""
	}

	uvIndex := weather.Daily.UvIndexMax[0]
	rec := ""

	if uvIndex >= uvIndexExtraLevel {
		rec = "Солнце сегодня очень мощное - обязательно нанесите SPF 30+, наденьте головной убор и избегайте прямых лучей."
	} else if uvIndex >= uvIndexMediumLevel {
		rec = "Солнце сегодня в ударе - рекомендуем нанести SPF 15+ и надеть головной убор."
	}

	return rec
}
