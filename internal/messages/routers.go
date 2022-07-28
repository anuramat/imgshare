package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func MessageRouter(msg *tgbotapi.Message, userID int64, users models.Users, images models.Images) models.ChattableSlice {
	return models.ChattableSlice{MessageRouterSingle(msg, userID, users, images)}
}

func MessageRouterSingle(msg *tgbotapi.Message, userID int64, users models.Users, images models.Images) tgbotapi.Chattable {
	chatID := msg.Chat.ID
	switch users[userID].State {
	case models.StartState:
		return StartHandler(chatID, userID, users)
	case models.UploadImageInitState:
		return UploadImageInitHandler(chatID, userID, users)
	case models.UploadImageState:
		return UploadImageHandler(chatID, userID, users, msg)
	case models.UploadDescriptionState:
		return UploadDescriptionHandler(chatID, userID, users, images, msg.Text)
	case models.EditDescriptionState:
		return EditDescriptionHandler(chatID, userID, users, images, msg.Text)
	case models.RandomImageState:
		return RandomImageHandler(chatID, userID, images, users)
	case models.GalleryState:
		return GalleryHandler(chatID, userID, images, users)
	case models.NoState:
		return DefaultHandler(chatID, userID)
	}
	panic("Unreachable, check if all states are covered")
}
