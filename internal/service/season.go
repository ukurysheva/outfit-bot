package service

import (
	"time"

	"outfitbot/internal/model"
)

var monthToSeason = map[time.Month]model.Season{
	time.January:   model.SeasonWinter,
	time.February:  model.SeasonWinter,
	time.March:     model.SeasonSpring,
	time.April:     model.SeasonSpring,
	time.May:       model.SeasonSpring,
	time.June:      model.SeasonSummer,
	time.July:      model.SeasonSummer,
	time.August:    model.SeasonSummer,
	time.September: model.SeasonAutumn,
	time.October:   model.SeasonAutumn,
	time.November:  model.SeasonAutumn,
	time.December:  model.SeasonWinter,
}

func (s *Service) getCurrentSeason() model.Season {
	currentMonth := time.Now().Month()

	return monthToSeason[currentMonth]
}
