package main

import (
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/imgsharebot"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting bot...")

	token := os.Getenv("TELEGRAM_APITOKEN")
	if len(token) == 0 {
		log.Fatal("Telegram API token not found, exiting.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Couldn't initialize bot:", err)
	}
	server_address := os.Getenv("SERVER")
	if len(server_address) == 0 {
		log.Fatal("No imgshare server address in environment, exiting.")
	}

	conn, err := grpc.Dial(server_address)

	if err != nil {
		log.Fatal("Can't dial imgshare server, exiting.")
	}
	defer conn.Close()
	client := api.NewImgShareClient(conn)
	// start main loop
	data := &models.BotData{}
	data.Client = client
	data.Users = models.Users{}
	data.MessageFiles = models.MessageFiles{}
	ctx := context.Background()
	imgsharebot.StartBot(ctx, bot, data)
}
