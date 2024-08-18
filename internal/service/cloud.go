package service

import "outfitbot/internal/model"

var (
	cloudFull       = 85
	cloudAlmostFull = 55
	cloudLittle     = 25
	cloudAbsent     = 0

	cloud = []int{cloudFull, cloudAlmostFull, cloudLittle, cloudAbsent}

	cloudEmoji = map[int]string{
		cloudFull:       "☁️",
		cloudAlmostFull: "⛅️",
		cloudLittle:     "🌤",
		cloudAbsent:     "☀️",
	}
)

func (s *Service) CloudEmoji(weather *model.Weather) string {
	for _, v := range cloud {
		if weather.Current.CloudCover >= v {
			return cloudEmoji[v]
		}
	}

	return ""
}
