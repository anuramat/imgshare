package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func DefaultHandler(chatID int64, userID int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, "Choose a command from the menu.")
}

func StartHandler(chatID int64, userID int64, users models.Users) tgbotapi.MessageConfig {
	users[userID].State = models.NoState
	return tgbotapi.NewMessage(chatID, "Welcome! I guess it would be nice if I told you what this bot does, but I'm too sleep deprived.")
}

func UploadImageInitHandler(chatID int64, userID int64, users models.Users) tgbotapi.MessageConfig {
	users[userID].State = models.UploadImageState
	return tgbotapi.NewMessage(chatID, "Send me an image!")
}

func UploadImageHandler(chatID int64, userID int64, users models.Users, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	if len(msg.Photo) == 0 {
		return tgbotapi.NewMessage(chatID, "Invalid image, try again.")
	}
	fid := string(msg.Photo[len(msg.Photo)-1].FileID)
	users[userID].LastUpload = fid
	users[userID].State = models.UploadDescriptionState
	return tgbotapi.NewMessage(chatID, "Send the description for your image")
}

func UploadDescriptionHandler(chatID int64, userID int64, users models.Users, images models.Images, description string) tgbotapi.MessageConfig {
	fid := users[userID].LastUpload
	users[userID].Images = append(users[userID].Images, fid)
	images[fid] = &(models.ImageMeta{Description: description, Votes: make(map[int64]bool)})
	users[userID].State = models.NoState
	return tgbotapi.NewMessage(chatID, "Uploaded successfully!")
}

func EditDescriptionHandler(chatID int64, userID int64, users models.Users, images models.Images, description string) tgbotapi.MessageConfig {
	fid := users[userID].LastUpload
	images[fid].Description = description
	users[userID].State = models.NoState
	return tgbotapi.NewMessage(chatID, "Description changed successfully!")
}

func RandomImageHandler(chatID int64, userID int64, images models.Images, users models.Users) tgbotapi.Chattable {
	if len(images) == 0 {
		return tgbotapi.NewMessage(chatID, "No images yet...")
	}
	fileID := models.GetRandomImage(images)
	text := models.PublicImageText(fileID, images)
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(fileID))
	photo.Caption = text
	photo.ReplyMarkup = keyboards.PublicImageKeyboard
	users[userID].LastDownload = fileID
	return photo
}

func GalleryHandler(chatID int64, userID int64, images models.Images, users models.Users) tgbotapi.Chattable {
	u := users[userID]
	n_photos := len(u.Images)
	if n_photos == 0 {
		return tgbotapi.NewMessage(chatID, "You don't have images yet...")
	}
	fileID := u.Images[0]
	users[userID].LastGalleryIndex = 0
	text := models.GalleryText(0, fileID, userID, images, users)
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(fileID))
	photo.Caption = text
	photo.ReplyMarkup = keyboards.GalleryKeyboard
	users[userID].LastDownload = fileID
	return photo
}
