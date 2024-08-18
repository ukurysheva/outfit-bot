package service

import (
	"fmt"
	"math"

	"outfitbot/internal/model"
)

var (
	temperatureTmpl = "%s Температура %d°C, ощущается как %d°C."

	tempPlus            = []int{tempPlusHot, tempPlusAlmostHot, tempPlusComfy, tempPlusAlmostComfy, tempPlusLittleCold, tempPlusCold, tempPlusExtraCold}
	tempPlusHot         = 25
	tempPlusAlmostHot   = 22
	tempPlusComfy       = 19
	tempPlusAlmostComfy = 16
	tempPlusLittleCold  = 12
	tempPlusCold        = 7
	tempPlusExtraCold   = 0
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
