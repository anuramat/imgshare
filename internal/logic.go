package internal

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/callbacks"
	"gitlab.ozon.dev/anuramat/homework-1/internal/commands"
	"gitlab.ozon.dev/anuramat/homework-1/internal/messages"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func StartBot(ctx context.Context, bot *tgbotapi.BotAPI, data *models.BotData) {
	// offset magic
	updateConfig := tgbotapi.NewUpdate(0)
	// long polling timeout (seconds)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	var chattableSlice models.ChattableSlice
	var uid int64
	for update := range updates {
		if update.Message != nil {
			uid = update.Message.From.ID
			// make sure user is in storage
			if _, ok := data.Users[uid]; !ok {
				data.AddUser(uid)
			}

			// check if the message is a command
			var cmd_err error
			if update.Message.IsCommand() {
				cmd_err = commands.CommandRouter(update.Message.Command(), uid, data)
			}

			// construct msg
			if cmd_err != nil {
				chattableSlice = models.ChattableSlice{tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid command")}
			} else {
				chattableSlice = messages.MessageRouter(ctx, update.Message, uid, data)
			}
		} else if update.CallbackQuery != nil {
			chattableSlice = callbacks.CallbackRouter(ctx, update.CallbackQuery, data)
		} else {
			continue
		}

		// sending a response
		for _, chattable := range chattableSlice {
			switch chattable.(type) {
			case tgbotapi.DeleteMessageConfig:
				_, err := bot.Request(chattable)
				if err != nil {
					log.Panicln("Encountered an error when sending a message:", err)
				}
			default:
				sent, err := bot.Send(chattable)
				if err != nil {
					log.Panicln("Encountered an error when sending a message:", err)
				}
				// HACK
				if len(sent.Photo) != 0 {
					data.MessageFiles[int64(sent.MessageID)] = data.Users[uid].LastDownload
				}
			}
		}

	}
}
