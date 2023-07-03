package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Start    = "/start"
	Add      = "Add"
	Remove   = "remove"
	ShowNote = "Show Note"
)

func GetMyKeyBoards() tgbotapi.ReplyKeyboardMarkup {

	var numericKeyboard = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Add"),
			tgbotapi.NewKeyboardButton("Show Note"),
		),
	)
	return numericKeyboard
}

var NumericKeyboard2 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Удалить", "KEY11"),
	),
)
var Num3kb = tgbotapi.NewInlineKeyboardButtonData("Удалить", "KEY11")
