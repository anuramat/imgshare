package callbacks

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func upvoteCallback(ctx context.Context, userID int64, fileID string, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	req := api.ImageAuthRequest{Image: &api.Image{FileID: fileID}, UserID: userID}
	image, err := data.Client.UpvoteImage(ctx, &req)
	if err != nil {
		return models.ChattableSlice{}
		// TODO return error message
	}
	text := models.PublicImageText(int(image.Upvotes), int(image.Downvotes), image.Description)
	changeDescription := tgbotapi.NewEditMessageCaption(chatID, messageID, text)
	changeDescription.BaseEdit.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeDescription}
}

func downvoteCallback(ctx context.Context, userID int64, fileID string, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	req := api.ImageAuthRequest{Image: &api.Image{FileID: fileID}, UserID: userID}
	image, err := data.Client.DownvoteImage(ctx, &req)
	if err != nil {
		return models.ChattableSlice{}
		// TODO return error message
	}
	text := models.PublicImageText(int(image.Upvotes), int(image.Downvotes), image.Description)
	changeDescription := tgbotapi.NewEditMessageCaption(chatID, messageID, text)
	changeDescription.BaseEdit.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeDescription}
}

func editDescriptionCallback(userID int64, fileID string, chatID int64, data *models.BotData) models.ChattableSlice {
	data.Users[userID].LastUpload = fileID
	data.Users[userID].State = models.EditDescriptionState
	return models.ChattableSlice{tgbotapi.NewMessage(chatID, "Enter new description:")}
}

// TODO mvoe more logic to indexImage
func nextImageCallback(ctx context.Context, userID int64, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	index := data.Users[userID].LastGalleryIndex + 1
	if index < 0 {
		index = 0
	}
	req := api.GalleryRequest{Offset: int32(index), UserID: userID}
	result, err := data.Client.GetGalleryImage(ctx, &req)
	// TODO return err if not nil
	if err != nil {
		log.Panicln(err)
	}
	data.Users[userID].LastGalleryIndex = int(result.Offset) + 1
	return indexImage(index, int(result.Total), result.Image, userID, chatID, messageID, data)
}

func previousImageCallback(ctx context.Context, userID int64, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	index := data.Users[userID].LastGalleryIndex - 1
	if index < 0 {
		index = 0
	}
	req := api.GalleryRequest{Offset: int32(index), UserID: userID}
	result, err := data.Client.GetGalleryImage(ctx, &req)
	// TODO return err if not nil
	if err != nil {
		log.Panicln(err)
	}
	data.Users[userID].LastGalleryIndex = int(result.Offset) + 1
	return indexImage(index, int(result.Total), result.Image, userID, chatID, messageID, data)
}

func indexImage(index, n_photos int, image *api.Image, userID int64, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	fileID := image.FileID
	changeImage := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &keyboards.GalleryKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(fileID)),
	}
	caption := models.GalleryText(index, n_photos, int(image.Upvotes), int(image.Downvotes), image.Description)
	changeText := tgbotapi.NewEditMessageCaption(chatID, messageID, caption)
	changeText.ReplyMarkup = &keyboards.GalleryKeyboard
	return models.ChattableSlice{changeImage, changeText}
}

func deleteImageCallback(ctx context.Context, userID int64, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	// TODO
	return models.ChattableSlice{tgbotapi.NewDeleteMessage(chatID, messageID)}
	// index := users[userID].LastGalleryIndex
	// // TODO return image with same index
	// users[userID].State = models.NoState
	// return models.ChattableSlice{tgbotapi.NewDeleteMessage(chatID, messageID)}
}

func randomImageCallback(ctx context.Context, userID int64, chatID int64, messageID int, data *models.BotData) models.ChattableSlice {
	image, err := data.Client.GetRandomImage(ctx, &api.Empty{})
	if err != nil {
		log.Panicln(err)
	}
	// TODO err
	data.Users[userID].LastDownload = image.FileID
	changeImage := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &keyboards.PublicImageKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(image.FileID)),
	}
	caption := models.PublicImageText(int(image.Upvotes), int(image.Downvotes), image.Description)
	changeText := tgbotapi.NewEditMessageCaption(chatID, messageID, caption)
	changeText.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeImage, changeText}
}
