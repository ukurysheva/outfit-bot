package service

import (
	"fmt"

	"outfitbot/internal/model"
)

var precipitationTmpl = "%s️ Вероятность осадков %d%%."

const (
	isRainyBarrier      = 30
	isSuperRainyBarrier = 80
)

func (s *Service) PrecipitationInfo(weather *model.Weather) string {
	return fmt.Sprintf(precipitationTmpl, s.getPrecipitationEmoji(weather.Daily.PrecipitationProbability[0]), weather.Daily.PrecipitationProbability[0])
}

func (s *Service) PrecipitationRecommendation(weather *model.Weather) string {
	if weather.Daily.PrecipitationProbability[0] > isRainyBarrier {
		return "Не забудьте зонт!"
	}

	return ""
}

func (s *Service) getPrecipitationEmoji(probability int) string {
	var emoji string
	season := s.getCurrentSeason()

	switch season {
	case model.SeasonSummer, model.SeasonAutumn, model.SeasonSpring:
		if probability > isSuperRainyBarrier {
			emoji = "💦"
		} else if probability > isRainyBarrier {
			emoji = "☔️"
		} else {
			emoji = "☂️"
		}
	case model.SeasonWinter:
		emoji = "☃️"
	}

	return emoji
}
