package callbacks

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func CallbackRouter(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	switch query.Data {
	case keyboards.UpvoteImageButton:
		return upvoteCallback(ctx, query, data)
	case keyboards.DownvoteImageButton:
		return downvoteCallback(ctx, query, data)
	case keyboards.NextImageButton:
		return nextImageCallback(ctx, query, data)
	case keyboards.PreviousImageButton:
		return previousImageCallback(ctx, query, data)
	case keyboards.EditDescriptionButton:
		return editDescriptionCallback(ctx, query, data)
	case keyboards.DeleteImageButton:
		return deleteImageCallback(ctx, query, data)
	case keyboards.RandomImageButton:
		return randomImageCallback(ctx, query, data)
	}
	panic("Unreachable, check if all keyboard buttons are covered")
}
