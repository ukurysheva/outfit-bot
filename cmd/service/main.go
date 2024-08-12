package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	"outfitbot/internal/handlers"
	openmeteo "outfitbot/internal/pkg/open-meteo"
	"outfitbot/internal/service"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("loading .env file failed: %v", err)
	}

	tgToken := os.Getenv("TG_TOKEN")
	if tgToken == "" {
		log.Fatalf("telegram token is not set")
	}

	botApi, err := tgbotapi.NewBotAPI(tgToken)
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
