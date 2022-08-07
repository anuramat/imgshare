package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gitlab.ozon.dev/anuramat/homework-1/internal"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func main() {
	log.Println("Starting bot...")

	err := godotenv.Load()
	if err == nil {
		log.Println("Reading .env file.")
	}

	token := os.Getenv("TELEGRAM_APITOKEN")
	if len(token) == 0 {
		log.Fatal("Telegram API token not found in environment, exiting.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Couldn't initialize bot:", err)
	}

	// start main loop
	users := models.Users{}
	images := models.Images{}
	messageFiles := models.MessageFiles{}
	internal.StartBot(bot, users, images, messageFiles)
}
