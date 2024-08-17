package service

import (
	"fmt"
	"math"

	"outfitbot/internal/model"
)

var (
	windLight    = 1
	windModerate = 5
	windHard     = 9
	windExtra    = 12

	windZeroTmpl     = "😌 Ветер отсутствует."
	windLightTmpl    = "%s Ветер слабый, около %d м/с."
	windModerateTmpl = "🪁 Ветер ощутимый, %d м/с."
	windHardTmpl     = "🌪 Ветер очень сильный, около %d м/с."
	windExtraTmpl    = "🌪 Сегодня на улице ураган - скорость %d м/с."
)

func (s *Service) WindInfo(weather *model.Weather) string {
	wind := int(math.Round(weather.Current.WindSpeed))

	if wind >= windExtra {
		return fmt.Sprintf(windExtraTmpl, wind)
	} else if wind >= windHard {
		return fmt.Sprintf(windHardTmpl, wind)
	} else if wind >= windModerate {
		return fmt.Sprintf(windModerateTmpl, wind)
	} else if wind >= windLight {
		return fmt.Sprintf(windLightTmpl, s.getWindEmoji(), wind)
	}

	return windZeroTmpl
}

func (s *Service) getWindEmoji() string {
	var emoji string
	season := s.getCurrentSeason()

	switch season {
	case model.SeasonSpring:
		emoji = "🌱"
	case model.SeasonSummer:
		emoji = "🍃"
	case model.SeasonAutumn:
		emoji = "🍂"
	case model.SeasonWinter:
		emoji = "🌬"

	}

	return emoji
}
