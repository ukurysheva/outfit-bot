package open_meteo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"outfitbot/internal/model"
)

var (
	host   = "https://api.open-meteo.com/v1/forecast?"
	params = "latitude=%.5f&longitude=%.2f&current=temperature_2m,apparent_temperature,precipitation,rain,weather_code,cloud_cover,wind_speed_10m&daily=weather_code,temperature_2m_max,temperature_2m_min,uv_index_max,precipitation_probability_max&forecast_days=1"
)

func (c *Client) GetCurrentWeather(lat, long float64) (*model.Weather, error) {
	resp, err := http.Get(host + fmt.Sprintf(params, lat, long))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned status code: %d", resp.StatusCode)
	}

	var currentWeather model.Weather
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body from resp: %w", err)
	}

	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal current weather resp: %w", err)
	}

	currentWeather.Current.WindSpeed = currentWeather.Current.WindSpeed * 1000 / 3600

	return &currentWeather, nil
}
