package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"telegram-bot/bot/structure"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("token")

	tgBot := structure.NewBot(&structure.TGBotConfig{
		Token:   token,
		TimeOut: 60,
		Debug:   true,
	})

	tgBot.Start()
}
