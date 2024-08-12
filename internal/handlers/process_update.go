package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"outfitbot/internal/model"
)

func (h *Handlers) ProcessUpdate(update tgbotapi.Update) {
	var user *model.User

	// keyboard callback
	if update.CallbackQuery != nil {
		user = h.getUserFromChat(update.CallbackQuery.Message.Chat)

		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		_, err := h.client.AnswerCallbackQuery(callback)
		if err != nil {
			fmt.Println(err)
		}

		h.processCallback(update, user)

		return
	}

	if update.Message != nil {
		user = h.getUserFromChat(update.Message.Chat)

		// command callback
		if update.Message.IsCommand() {
			h.processCommand(update, user)
		}
	}
}

func (h *Handlers) getUserFromChat(chat *tgbotapi.Chat) *model.User {
	if chat == nil {
		return nil
	}

	return &model.User{ChatID: chat.ID}
}
