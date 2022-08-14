package messages

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func MessageRouter(ctx context.Context, msg *tgbotapi.Message, userID int64, data *models.BotData) models.ChattableSlice {
	return models.ChattableSlice{MessageRouterSingle(ctx, msg, userID, data)}
}

func MessageRouterSingle(ctx context.Context, msg *tgbotapi.Message, userID int64, data *models.BotData) tgbotapi.Chattable {
	chatID := msg.Chat.ID
	switch data.Users[userID].State {
	case models.StartState:
		return StartHandler(chatID, userID, data)
	case models.UploadImageInitState:
		return UploadImageInitHandler(chatID, userID, data)
	case models.UploadImageState:
		return UploadImageHandler(chatID, userID, data, msg)
	case models.UploadDescriptionState:
		return UploadDescriptionHandler(ctx, chatID, userID, data, msg.Text)
	case models.EditDescriptionState:
		return EditDescriptionHandler(ctx, chatID, userID, data, msg.Text)
	case models.RandomImageState:
		return RandomImageHandler(ctx, chatID, userID, data)
	case models.GalleryState:
		return GalleryHandler(ctx, chatID, userID, data)
	case models.NoState:
		return DefaultHandler(chatID, userID)
	}
	panic("Unreachable, check if all states are covered")
}
