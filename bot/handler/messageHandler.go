package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MessageHandler struct {
}

func (m MessageHandler) Handle(mess *tgbotapi.Message) tgbotapi.MessageConfig {
	if mess.IsCommand() {
		return m.commandHandler(mess)
	}
	return tgbotapi.NewMessage(mess.Chat.ID, mess.Text)
}

func (m MessageHandler) commandHandler(mess *tgbotapi.Message) tgbotapi.MessageConfig {
	var response tgbotapi.MessageConfig
	switch mess.Command() {
	case "help":
		response = tgbotapi.NewMessage(mess.Chat.ID, "Введена команда help")
	default:
		response = tgbotapi.NewMessage(mess.Chat.ID, "Введена неизвестная команда")
	}

	return response
}
