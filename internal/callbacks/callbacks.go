package callbacks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func upvoteCallback(userID int64, fileID string, chatID int64, messageID int, images models.Images, users models.Users) models.ChattableSlice {
	ok := models.UpvoteImage(images, userID, fileID)
	if !ok {
		return models.ChattableSlice{}
	}
	text := models.PublicImageText(fileID, images)
	changeDescription := tgbotapi.NewEditMessageCaption(chatID, messageID, text)
	changeDescription.BaseEdit.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeDescription}
}

func downvoteCallback(userID int64, fileID string, chatID int64, messageID int, images models.Images, users models.Users) models.ChattableSlice {
	ok := models.DownvoteImage(images, userID, fileID)
	if !ok {
		return models.ChattableSlice{}
	}
	text := models.PublicImageText(fileID, images)
	changeDescription := tgbotapi.NewEditMessageCaption(chatID, messageID, text)
	changeDescription.BaseEdit.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeDescription}
}

func editDescriptionCallback(userID int64, fileID string, chatID int64, users models.Users) models.ChattableSlice {
	users[userID].LastUpload = fileID
	users[userID].State = models.EditDescriptionState
	return models.ChattableSlice{tgbotapi.NewMessage(chatID, "Enter new description:")}
}

// TODO combine next/prev callbacks
func nextImageCallback(userID int64, chatID int64, messageID int, users models.Users, images models.Images) models.ChattableSlice {
	n_files := len(users[userID].Images)
	index := (users[userID].LastGalleryIndex + 1) % n_files
	if index < 0 {
		index += n_files
	}
	users[userID].LastGalleryIndex = index
	return indexImage(index, userID, chatID, messageID, users, images)
}

func previousImageCallback(userID int64, chatID int64, messageID int, users models.Users, images models.Images) models.ChattableSlice {
	n_files := len(users[userID].Images)
	index := (users[userID].LastGalleryIndex - 1) % n_files
	if index < 0 {
		index += n_files
	}
	users[userID].LastGalleryIndex = index
	return indexImage(index, userID, chatID, messageID, users, images)
}

func indexImage(index int, userID int64, chatID int64, messageID int, users models.Users, images models.Images) models.ChattableSlice {
	fileID := users[userID].Images[index]
	changeImage := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &keyboards.GalleryKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(fileID)),
	}
	description := models.GalleryText(index, fileID, userID, images, users)
	changeText := tgbotapi.NewEditMessageCaption(chatID, messageID, description)
	changeText.ReplyMarkup = &keyboards.GalleryKeyboard
	return models.ChattableSlice{changeImage, changeText}
}

func deleteImageCallback(userID int64, chatID int64, messageID int, users models.Users, images models.Images) models.ChattableSlice {
	i := users[userID].LastGalleryIndex
	fileID := users[userID].Images[i]
	users[userID].Images = append(users[userID].Images[:i], users[userID].Images[i+1:]...)
	users[userID].State = models.NoState
	delete(images, fileID)
	return models.ChattableSlice{tgbotapi.NewDeleteMessage(chatID, messageID)}
}

func randomImageCallback(userID int64, chatID int64, messageID int, users models.Users, images models.Images) models.ChattableSlice {
	fileID := models.GetRandomImage(images)
	users[userID].LastDownload = fileID
	changeImage := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &keyboards.PublicImageKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(fileID)),
	}
	description := models.PublicImageText(fileID, images)
	changeText := tgbotapi.NewEditMessageCaption(chatID, messageID, description)
	changeText.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeImage, changeText}
}
