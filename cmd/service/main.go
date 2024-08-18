package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"outfitbot/internal/handlers"
	openmeteo "outfitbot/internal/pkg/open-meteo"
	"outfitbot/internal/service"
)

const (
	tgTokenEnv = "TG_TOKEN"
)

func main() {
	botApi, err := tgbotapi.NewBotAPI(getTokenFromEnv())
	if err != nil {
		log.Fatalf("telegram create bot api failed: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := botApi.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	cl := openmeteo.NewClient()
	srv := service.NewService(cl)

	h := handlers.NewHandlers(botApi, srv)
	for update := range updates {
		h.ProcessUpdate(update)
	}
}

func getTokenFromEnv() string {
	tgToken := os.Getenv(tgTokenEnv)
	if tgToken == "" {
		log.Fatalf("telegram token is not set")
	}

	return tgToken
}
