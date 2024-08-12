package service

import (
	"errors"
	"log"

	"outfitbot/internal/model"
)

func (s *Service) GetRecommendation(city string) (string, error) {
	cityCords, ex := model.CityCordsByName[city]
	if !ex {
		log.Printf("no coordinates found for city: %s", city)

		return "", errors.New("no coordinates found for city")
	}

	cityTitle, ex := model.CityTitleByName[city]
	if !ex {
		log.Printf("no title found for city: %s", city)

		return "", errors.New("no title found for city")
	}

	weather, err := s.weatherAPI.GetCurrentWeather(cityCords[0], cityCords[1])
	if err != nil {
		log.Printf("failed to GetCurrentWeather: %v", err)

		return "", err
	}

	return s.GenerateWeatherMsg(cityTitle, weather), nil
}

func (s *Service) GenerateWeatherMsg(cityTitle string, weather *model.Weather) string {
	msg := s.CityInfo(cityTitle)
	msg += "\n\n" + s.TempInfo(weather)
	msg += "\n\n" + s.WindInfo(weather)

	if len(weather.Daily.PrecipitationProbability) > 0 {
		msg += "\n\n" + s.PrecipitationInfo(weather)
	}

	msg += "\n\n" + s.ClothesRecommendation(weather)

	if precipitationRec := s.PrecipitationRecommendation(weather); precipitationRec != "" {
		msg += "\n\n" + precipitationRec
	}

	return msg
}
