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
	switch data.Users[userID].State {
	case models.StartState:
		return StartHandler(ctx, msg, data)
	case models.UploadImageInitState:
		return UploadImageInitHandler(ctx, msg, data)
	case models.UploadImageState:
		return UploadImageHandler(ctx, msg, data)
	case models.UploadDescriptionState:
		return UploadDescriptionHandler(ctx, msg, data)
	case models.EditDescriptionState:
		return EditDescriptionHandler(ctx, msg, data)
	case models.RandomImageState:
		return RandomImageHandler(ctx, msg, data)
	case models.GalleryState:
		return GalleryHandler(ctx, msg, data)
	case models.NoState:
		return DefaultHandler(ctx, msg, data)
	}
	panic("Unreachable, check if all states are covered")
}
