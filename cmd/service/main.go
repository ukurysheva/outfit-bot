package main

import (
	"flag"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	"outfitbot/internal/handlers"
	openmeteo "outfitbot/internal/pkg/open-meteo"
	"outfitbot/internal/service"
)

const (
	tgTokenEnv    = "TG_TOKEN"
	tgTokenDevEnv = "TG_TOKEN_DEV"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("loading .env file failed: %v", err)
	}

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
	devArg := flag.Bool("dev", false, "dev")
	flag.Parse()

	tgTokenEnvName := tgTokenEnv
	if devArg != nil && *devArg {
		tgTokenEnvName = tgTokenDevEnv
	}

	tgToken := os.Getenv(tgTokenEnvName)
	if tgToken == "" {
		log.Fatalf("telegram token is not set")
	}

	return tgToken
}
