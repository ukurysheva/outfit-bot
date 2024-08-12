package service

import (
	"fmt"
	"math"

	"outfitbot/internal/model"
)

var (
	temperatureTmpl = "%s Температура %d°C, ощущается как %d°C."
	//tempPlusEmoji    = map[int]string{
	//	tempPlusHot:         "🥵",
	//	tempPlusAlmostHot:   "🌡️",
	//	tempPlusComfy:       "✨",
	//	tempPlusAlmostComfy: "🌤",
	//	tempPlusLittleCold:  "🕊️",
	//	tempPlusCold:        "🌬",
	//	tempPlusExtraCold:   "🥶",
	//}

	tempPlus            = []int{tempPlusHot, tempPlusAlmostHot, tempPlusComfy, tempPlusAlmostComfy, tempPlusLittleCold, tempPlusCold, tempPlusExtraCold}
	tempPlusHot         = 25
	tempPlusAlmostHot   = 22
	tempPlusComfy       = 19
	tempPlusAlmostComfy = 16
	tempPlusLittleCold  = 12
	tempPlusCold        = 7
	tempPlusExtraCold   = 0

	tempPlusRec = map[int][]string{
		tempPlusHot:         {"легкое льяное платье", "шорты с легкой майкой/топом", "легкие юбку/брюки и рубашку с коротким рукавом"},
		tempPlusAlmostHot:   {"легкие брюки и футболку", "юбку и легкую рубашку", "легкое платье с длинным рукавом"},
		tempPlusComfy:       {"джинсы, футболку и рубашку", "брюки и лонгслив", "платье с длинным рукавом"},
		tempPlusAlmostComfy: {"джинсы и толстовку", "брюки, футболку и кожаную куртку", "платье и легкую куртку"},
		tempPlusLittleCold:  {"джинсы, лонгслив и легкую куртку", "брюки, толстовку и плащ", "юбку, лонгслив/рубашку и пальто"},
		tempPlusCold:        {"брюки, лонгслив и пальто", "джинсы, толстовку и весеннюю куртку", "брюки, свитер и плащ"},
	}

	tempPlusAccessories = map[int]string{
		tempPlusHot:       "Не забудьте головной убор - кепку, платок и или панамку.",
		tempPlusExtraCold: "Лучше надеть шарф и, например, повязку на голову.",
		tempPlusCold:      "Лучше надеть шарф и, например, повязку на голову.",
	}
)

func (s *Service) TempInfo(weather *model.Weather) string {
	temp := int(math.Round(weather.Current.Temperature))
	apparentTemp := int(math.Round(weather.Current.ApparentTemperature))
	emoji := s.CloudEmoji(weather)

	return fmt.Sprintf(temperatureTmpl, emoji, temp, apparentTemp)
}

//
//func (s *Service) tempMinusRec(temp int) string {
//	var word string
//
//	//if temp >= tempExtra {
//	//	return fmt.Sprintf(tempExtraTmpl, temp)
//	//} else if temp >= tempHard {
//	//	word = tempHardWords[rand.Intn(len(tempHardWords)-1)]
//	//	return fmt.Sprintf(tempHardTmpl, word, temp)
//	//} else if temp >= tempModerate {
//	//	word = tempModerateWords[rand.Intn(len(tempModerateWords)-1)]
//	//	return fmt.Sprintf(tempModerateTmpl, word, temp)
//	//} else if temp >= tempLight {
//	//	word = tempLightWords[rand.Intn(len(tempLightWords)-1)]
//	//	return fmt.Sprintf(tempLightTmpl, word, wind)
//	//}
//
//	//word = windZeroWords[rand.Intn(len(windZeroWords)-1)]
//
//	return fmt.Sprintf(windZeroTmpl)
//}
