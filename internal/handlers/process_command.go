package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"outfitbot/internal/model"
)

func (h *Handlers) processCommand(update tgbotapi.Update, user *model.User) {
	var err error

	switch update.Message.Command() {
	case model.CommandStart:
		err = h.sendStartMsg(user)
	}

	if err != nil {
		h.sendSorryMsg(user)
	}
}
