package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ukurysheva/tglogger"

	"outfitbot/internal/model"
)

var (
	startMsg = "Приветы!\n" +
		"На связи Outfit Bot 💅🏻, и я помогу вам выбрать, что надеть сегодня, на основе погоды в вашем городе.\n\n" +
		"Пожалуйста, выберите свой город из списка ниже."

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
		tglogger.WithFields(tglogger.Fields{"user": user}).Errorf("failed to sendStartMsg: %v", err)

		return err
	}

	return nil
}

func (h *Handlers) sendChangeCityMsg(user *model.User, _ string) error {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(h.getCityButtons()...)

	if err := h.sendMsg(user, changeCityMsg, &keyboard); err != nil {
		tglogger.WithFields(tglogger.Fields{"user": user}).Errorf("failed to sendChangeCityMsg: %v", err)

		return err
	}

	return nil
}

func (h *Handlers) sendRecommendationMsg(user *model.User, city string) error {
	recMsg, err := h.rec.GetRecommendation(city)
	if err != nil {
		tglogger.WithFields(tglogger.Fields{"user": user, "city": city}).Errorf("failed to GetRecommendation: %v", err)

		return err
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(h.getContinueButtons(city))

	if err = h.sendMsg(user, recMsg, &keyboard); err != nil {
		tglogger.WithFields(tglogger.Fields{"user": user, "recMsg": recMsg}).Errorf("failed to sendRecommendationMsg: %v", err)

		return err
	}

	return nil
}

func (h *Handlers) sendSorryMsg(user *model.User) {
	if err := h.sendMsg(user, internalErrMsg, nil); err != nil {
		tglogger.WithFields(tglogger.Fields{"user": user}).Errorf("failed to sendSorryMsg: %v", err)
	}
}
