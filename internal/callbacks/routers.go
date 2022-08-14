package callbacks

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func CallbackRouter(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID
	fileID := data.MessageFiles[int64(messageID)]
	switch query.Data {
	case keyboards.UpvoteImageButton:
		return upvoteCallback(ctx, userID, fileID, chatID, messageID, data)
	case keyboards.DownvoteImageButton:
		return downvoteCallback(ctx, userID, fileID, chatID, messageID, data)
	case keyboards.NextImageButton:
		return nextImageCallback(ctx, userID, chatID, messageID, data)
	case keyboards.PreviousImageButton:
		return previousImageCallback(ctx, userID, chatID, messageID, data)
	case keyboards.EditDescriptionButton:
		return editDescriptionCallback(userID, fileID, chatID, data)
	case keyboards.DeleteImageButton:
		return deleteImageCallback(ctx, userID, chatID, messageID, data)
	case keyboards.RandomImageButton:
		return randomImageCallback(ctx, userID, chatID, messageID, data)
	}
	panic("Unreachable, check if all keyboard buttons are covered")
}
