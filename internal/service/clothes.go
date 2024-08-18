package service

import (
	"math"
	"math/rand"

	"outfitbot/internal/model"
)

const (
	maxRecommendationsCount = 3
)

var (
	beginWords = []string{"Рекомендуем надеть", "Можно надеть", "Рекомендуем", "Как вариант, можно надеть", "Что надеть"}

	tempPlusRec = map[int][]string{
		tempPlusHot:         {"легкие шорты и рубашку с коротким рукавом", "легкое льяное платье", "шорты с майкой/топом"},
		tempPlusAlmostHot:   {"джинсы и футболку", "брюки и легкую рубашку", "брюки и рубашку с коротким рукавом", "шорты и футболку, сверху рубашку", "брюки и майку, поверх - легкую рубашку", "легкое платье"},
		tempPlusComfy:       {"джинсы, футболку и рубашку", "брюки и лонгслив", "платье с длинным рукавом", "брюки и рубашка", "юбка и свитшот"},
		tempPlusAlmostComfy: {"джинсы и толстовку", "брюки, футболку и кожаную куртку", "костюм - брюки и пиджак", "платье и свитшот", "джинсы, футболку и тренч", "джинсы, рубашку и джемпер"},
		tempPlusLittleCold:  {"джинсы, лонгслив и легкую куртку", "джинсы, толстовку и плащ", "брюки, рубашка и джемпер", "костюм - брюки и пиджак, сверху плащ", "брюки и свитер", "юбку и рубашку, легкое пальто"},
		tempPlusCold:        {"брюки, лонгслив и пальто", "юбку, рубашку и пальто", "брюки, рубашку и легкую куртку", "джинсы, толстовку и легкую куртку", "брюки, свитер и плащ"},
		tempPlusExtraCold:   {"джинсы, толстовка и пальто", "джинсы, свитер и пальто", "юбка, свитер и пальто", "брюки, лонгслив и весенняя куртка", "брюки, рубашка и джемпер, легкая куртка", "брюки, свитер и легкая куртка"},
	}

	tempPlusAccessories = map[int]string{
		tempPlusHot:       "Не забудьте головной убор - кепку, платок и или панамку.",
		tempPlusExtraCold: "Лучше надеть шарф и, например, повязку на голову.",
		tempPlusCold:      "Лучше надеть шарф и, например, повязку на голову.",
	}
)

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
	for i, v := range recList {
		rec += "• " + v + "\n"

		if i+1 == maxRecommendationsCount {
			break
		}
	}

	if recAccessories, ok := tempPlusAccessories[currentBarier]; ok {
		rec += "\n" + recAccessories
	}

	return rec
}
