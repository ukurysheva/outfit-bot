package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"outfitbot/internal/model"
)

func (h *Handlers) getCityButtons() [][]tgbotapi.InlineKeyboardButton {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	row := make([]tgbotapi.InlineKeyboardButton, 2)

	i := 0
	for _, code := range model.CityNameList {
		row[i] = tgbotapi.NewInlineKeyboardButtonData(model.CityTitleByName[code], code)
		i++

		if i == 2 {
			i = 0
			rows = append(rows, row)
			row = make([]tgbotapi.InlineKeyboardButton, 2)
		}
	}

	return rows
}

func (h *Handlers) getContinueButtons(cityName string) []tgbotapi.InlineKeyboardButton {
	return []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData(model.RefreshButtonTitle, cityName),
		tgbotapi.NewInlineKeyboardButtonData(model.ChangeCityButtonTitle, model.ChangeCityButtonName),
	}
}
