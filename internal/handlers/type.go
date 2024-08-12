package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handlers struct {
	client *tgbotapi.BotAPI
	rec    RecommendService
}

func NewHandlers(bot *tgbotapi.BotAPI, rec RecommendService) *Handlers {
	return &Handlers{client: bot, rec: rec}
}

type RecommendService interface {
	GetRecommendation(city string) (string, error)
}
