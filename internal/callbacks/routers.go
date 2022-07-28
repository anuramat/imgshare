package callbacks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func CallbackRouter(query *tgbotapi.CallbackQuery, users models.Users, images models.Images, messageFiles models.MessageFiles) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID
	fileID := messageFiles[int64(messageID)]
	switch query.Data {
	case keyboards.UpvoteImageButton:
		return upvoteCallback(userID, fileID, chatID, messageID, images, users)
	case keyboards.DownvoteImageButton:
		return downvoteCallback(userID, fileID, chatID, messageID, images, users)
	case keyboards.NextImageButton:
		return nextImageCallback(userID, chatID, messageID, users, images)
	case keyboards.PreviousImageButton:
		return previousImageCallback(userID, chatID, messageID, users, images)
	case keyboards.EditDescriptionButton:
		return editDescriptionCallback(userID, fileID, chatID, users)
	case keyboards.DeleteImageButton:
		return deleteImageCallback(userID, chatID, messageID, users, images)
	case keyboards.RandomImageButton:
		return randomImageCallback(userID, chatID, messageID, users, images)
	}
	panic("Unreachable, check if all keyboard buttons are covered")
}
