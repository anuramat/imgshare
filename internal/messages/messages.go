package messages

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func DefaultHandler(chatID int64, userID int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, "Choose a command from the menu.")
}

func StartHandler(chatID int64, userID int64, data *models.BotData) tgbotapi.MessageConfig {
	data.Users[userID].State = models.NoState
	return tgbotapi.NewMessage(chatID, "Welcome! I guess it would be nice if I told you what this bot does, but I'm too sleep deprived.")
}

func UploadImageInitHandler(chatID int64, userID int64, data *models.BotData) tgbotapi.MessageConfig {
	data.Users[userID].State = models.UploadImageState
	return tgbotapi.NewMessage(chatID, "Send me an image!")
}

func UploadImageHandler(chatID int64, userID int64, data *models.BotData, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	if len(msg.Photo) == 0 {
		return tgbotapi.NewMessage(chatID, "Invalid image, try again.")
	}
	fid := string(msg.Photo[len(msg.Photo)-1].FileID)
	data.Users[userID].LastUpload = fid
	data.Users[userID].State = models.UploadDescriptionState
	return tgbotapi.NewMessage(chatID, "Send the description for your image")
}

func UploadDescriptionHandler(ctx context.Context, chatID int64, userID int64, data *models.BotData, description string) tgbotapi.MessageConfig {
	fid := data.Users[userID].LastUpload
	req := api.ImageAuthRequest{UserID: userID, Image: &api.Image{FileID: fid, Description: description}}
	_, err := data.Client.SetDescriptionImage(ctx, &req)
	if err != nil {
		log.Panicln(err)
	}
	data.Users[userID].State = models.NoState
	return tgbotapi.NewMessage(chatID, "Uploaded successfully!")
}

func EditDescriptionHandler(ctx context.Context, chatID int64, userID int64, data *models.BotData, description string) tgbotapi.MessageConfig {
	fid := data.Users[userID].LastUpload
	req := api.ImageAuthRequest{UserID: userID, Image: &api.Image{FileID: fid}}
	_, err := data.Client.SetDescriptionImage(ctx, &req)
	if err != nil {
		log.Panicln(err)
	}
	data.Users[userID].State = models.NoState
	return tgbotapi.NewMessage(chatID, "Description changed successfully!")
}

func RandomImageHandler(ctx context.Context, chatID int64, userID int64, data *models.BotData) tgbotapi.Chattable {
	result, err := data.Client.GetRandomImage(ctx, &api.Empty{})
	// TODO return err if not nil, logic for case where user doesn't have images yet
	if err != nil {
		log.Panicln(err)
	}

	text := models.PublicImageText(int(result.Upvotes), int(result.Downvotes), result.Description)
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(result.FileID))
	photo.Caption = text
	photo.ReplyMarkup = keyboards.PublicImageKeyboard
	data.Users[userID].LastDownload = result.FileID
	return photo
}

func GalleryHandler(ctx context.Context, chatID int64, userID int64, data *models.BotData) tgbotapi.Chattable {
	index := data.Users[userID].LastGalleryIndex
	req := api.GalleryRequest{Offset: int32(index), UserID: userID}
	result, err := data.Client.GetGalleryImage(ctx, &req)
	// TODO return err if not nil, logic for case where user doesn't have images yet
	if err != nil {
		log.Panicln(err)
	}

	text := models.GalleryText(index, int(result.Total), int(result.Image.Upvotes), int(result.Image.Downvotes), result.Image.Description)
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(result.Image.FileID))
	photo.Caption = text
	photo.ReplyMarkup = keyboards.GalleryKeyboard
	data.Users[userID].LastDownload = result.Image.FileID
	return photo
}
