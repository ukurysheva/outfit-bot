package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"outfitbot/internal/model"
)

func (h *Handlers) processCallback(update tgbotapi.Update, user *model.User) {
	if update.CallbackQuery.Message == nil || user == nil {
		return
	}

	buttonCode := update.CallbackQuery.Data
	if err := h.callHandlerFunc(user, buttonCode); err != nil {
		h.sendSorryMsg(user)
	}
}

func (h *Handlers) callHandlerFunc(user *model.User, buttonCode string) error {
	var err error

	switch {
	case model.IsCityName(buttonCode):
		err = h.sendRecommendationMsg(user, buttonCode)
	case model.IsChangeCityButton(buttonCode):
		err = h.sendChangeCityMsg(user, buttonCode)
	}

	if err != nil {
		return err
	}

	return nil
}
