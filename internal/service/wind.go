package service

import (
	"fmt"
	"math"

	"outfitbot/internal/model"
)

var (
	// windZeroWords     = []string{"прям совсем по нулям", "кайфуем", "можно спокойно идти с распущенными волосами"}
	// windLightWords    = []string{"слабенький", "почти не чувствуется", "лайтовый", "практически отсутствует"}
	// windModerateWords = []string{"нормально так поднялся", "нормальный такой", "ощущается нормально так", "достаточно неприятный"}
	// windHardWords     = []string{"жестко", "пипец", "лютейше дует"}

	windLight    = 1
	windModerate = 5
	windHard     = 9
	windExtra    = 12

	windZeroTmpl     = "💨 Ветер отсутствует."
	windLightTmpl    = "💨 Ветер слабый, около %d м/с."
	windModerateTmpl = "💨 Ветер ощутимый, %d м/с."
	windHardTmpl     = "🌪 Ветер очень сильный, около %d м/с."
	windExtraTmpl    = "🌪 Сегодня на улице ураган - скорость %d м/с."
)

func (s *Service) WindInfo(weather *model.Weather) string {
	wind := int(math.Round(weather.Current.WindSpeed))

	if wind >= windExtra {
		return fmt.Sprintf(windExtraTmpl, wind)
	} else if wind >= windHard {
		// word = windHardWords[rand.Intn(len(windHardWords)-1)]
		return fmt.Sprintf(windHardTmpl, wind)
	} else if wind >= windModerate {
		// word = windModerateWords[rand.Intn(len(windModerateWords)-1)]
		return fmt.Sprintf(windModerateTmpl, wind)
	} else if wind >= windLight {
		// word = windLightWords[rand.Intn(len(windLightWords)-1)]
		return fmt.Sprintf(windLightTmpl, wind)
	}

	return windZeroTmpl
}
