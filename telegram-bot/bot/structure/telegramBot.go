package structure

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"telegram-bot/bot/handler"
)

type TGBotConfig struct {
	Token   string
	TimeOut int
	Debug   bool
}

type TGBot struct {
	Config  *TGBotConfig
	Handler handler.MessageHandler
	BotApi  *tgbotapi.BotAPI
}

func NewBot(config *TGBotConfig) *TGBot {
	botApi, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = config.Debug
	return &TGBot{
		Config: config,
		BotApi: botApi,
	}
}

func (b *TGBot) Start() {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = b.Config.TimeOut
	updates := b.BotApi.GetUpdatesChan(update)

	log.Printf("Bot [ %s ] was started", b.BotApi.Self.UserName)
	b.update(updates)
}

func (b *TGBot) update(updates tgbotapi.UpdatesChannel) {
	for updates := range updates {
		go func(updates tgbotapi.Update) {
			message := updates.Message
			if message == nil {
				return
			}

			msg := b.Handler.Handle(message)
			if _, err := b.BotApi.Send(msg); err != nil {
				log.Fatal(err)
			}

		}(updates)
	}
}
