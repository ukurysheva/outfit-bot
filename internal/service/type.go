package service

import (
	"outfitbot/internal/model"
)

type Service struct {
	weatherAPI WeatherAPI
}

func NewService(api WeatherAPI) *Service {
	return &Service{weatherAPI: api}
}

type WeatherAPI interface {
	GetCurrentWeather(lat, long float64) (*model.Weather, error)
}
