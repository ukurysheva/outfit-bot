package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"outfitbot/internal/model"
)

var (
	startMsg = "Привет!\n" +
		"На связи Lets's wear Bot, и я предложу вам &#128133;outfit&#128133; на основе погоды в вашем городе.\n" +
		"Пожалуйста, выберите свой город из списка ниже"

	changeCityMsg = "Пожалуйста, выберите свой город из списка ниже"

	internalErrMsg = "Не удалось сформировать прогноз одежды для вас 😢 Пожалуйста, попробуйте позже"
)

func (h *Handlers) sendMsg(user *model.User, text string, keyboard *tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewMessage(user.ChatID, text)
	msg.ParseMode = "html"

	if user.ReplyTo > 0 {
		msg.ReplyToMessageID = user.ReplyTo
	}

	if keyboard != nil {
		msg.ReplyMarkup = *keyboard
	}

	_, err := h.client.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handlers) sendStartMsg(user *model.User) error {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(h.getCityButtons()...)

	if err := h.sendMsg(user, startMsg, &keyboard); err != nil {
		log.Printf("failed to sendStartMsg: %v", err)

		return err
	}

	return nil
}

func (h *Handlers) sendChangeCityMsg(user *model.User, _ string) error {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(h.getCityButtons()...)

	if err := h.sendMsg(user, changeCityMsg, &keyboard); err != nil {
		log.Printf("failed to sendRecommendationMsg: %v", err)

		return err
	}

	return nil
}

func (h *Handlers) sendRecommendationMsg(user *model.User, city string) error {
	recMsg, err := h.rec.GetRecommendation(city)
	if err != nil {
		log.Printf("failed to GetRecommendation: %v", err)

		return err
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(h.getContinueButtons(city))

	if err = h.sendMsg(user, recMsg, &keyboard); err != nil {
		log.Printf("failed to sendRecommendationMsg: %v", err)

		return err
	}

	return nil
}

func (h *Handlers) sendSorryMsg(user *model.User) {
	if err := h.sendMsg(user, internalErrMsg, nil); err != nil {
		log.Printf("failed to sendSorryMsg: %v", err)
	}
}
