package service

import (
	"math"
	"math/rand"

	"outfitbot/internal/model"
)

var beginWords = []string{"Рекомендуем надеть", "Можно надеть", "Рекомендуем", "Как вариант, можно надеть", "Что надеть"}

func (s *Service) ClothesRecommendation(weather *model.Weather) string {
	temp := int(math.Round(weather.Current.ApparentTemperature))

	beginWord := beginWords[rand.Intn(len(beginWords)-1)]
	rec := beginWord + ":\n"

	if temp >= 0 {
		rec += s.tempPlusRec(temp)
	}

	return rec
}

func (s *Service) tempPlusRec(temp int) string {
	var recList []string
	var currentBarier int
	for _, v := range tempPlus {
		if temp >= v {
			currentBarier = v
			recList = tempPlusRec[v]

			break
		}
	}

	rand.Shuffle(len(recList), func(i, j int) { recList[i], recList[j] = recList[j], recList[i] })

	var rec string
	for _, v := range recList {
		rec += "• " + v + "\n"
	}

	if recAccessories, ok := tempPlusAccessories[currentBarier]; ok {
		rec += "\n" + recAccessories
	}

	return rec
}
