package callbacks

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/keyboards"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
	"gitlab.ozon.dev/anuramat/homework-1/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func upvoteCallback(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID
	fileID := data.MessageFiles[int64(messageID)]

	req := api.ImageAuthRequest{Image: &api.Image{FileID: fileID}, UserID: userID}
	image, err := data.Client.UpvoteImage(ctx, &req)
	if err != nil {
		log.Println(err)
		return models.ChattableSlice{tgbotapi.NewMessage(chatID, "Error!")}
	}
	text := utils.PublicImageText(int(image.Upvotes), int(image.Downvotes), image.Description)
	changeDescription := tgbotapi.NewEditMessageCaption(chatID, messageID, text)
	changeDescription.BaseEdit.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeDescription}
}

func downvoteCallback(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID
	fileID := data.MessageFiles[int64(messageID)]

	req := api.ImageAuthRequest{Image: &api.Image{FileID: fileID}, UserID: userID}
	image, err := data.Client.DownvoteImage(ctx, &req)
	if err != nil {
		log.Println(err)
		return models.ChattableSlice{tgbotapi.NewMessage(chatID, "Error!")}
	}
	text := utils.PublicImageText(int(image.Upvotes), int(image.Downvotes), image.Description)
	changeDescription := tgbotapi.NewEditMessageCaption(chatID, messageID, text)
	changeDescription.BaseEdit.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeDescription}
}

func editDescriptionCallback(_ context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID
	fileID := data.MessageFiles[int64(messageID)]

	data.Users[userID].LastUpload = fileID
	data.Users[userID].State = models.EditDescriptionState
	return models.ChattableSlice{tgbotapi.NewMessage(chatID, "Enter new description:")}
}

func nextImageCallback(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	return deltaIndexImage(1, ctx, query, data)
}

func previousImageCallback(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	return deltaIndexImage(-1, ctx, query, data)
}

func deltaIndexImage(delta_index int, ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID

	user_index := data.Users[userID].LastGalleryIndex + delta_index
	if user_index < 0 {
		user_index = 0
	}
	req := api.GalleryRequest{Offset: int32(user_index), UserID: userID}
	result, err := data.Client.GetGalleryImage(ctx, &req)
	// if gallery empty : delete message
	st, ok := status.FromError(err)
	if !ok {
		log.Println(err)
		return models.ChattableSlice{}
	} else if st.Code() == codes.NotFound {
		return models.ChattableSlice{tgbotapi.NewDeleteMessage(chatID, messageID)}
	}
	result_index := int(result.Offset)
	data.Users[userID].LastGalleryIndex = result_index
	image := result.Image

	changeImage := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &keyboards.GalleryKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(image.FileID)),
	}
	caption := utils.GalleryText(result_index, int(result.Total), int(image.Upvotes), int(image.Downvotes), image.Description)
	changeText := tgbotapi.NewEditMessageCaption(chatID, messageID, caption)
	changeText.ReplyMarkup = &keyboards.GalleryKeyboard
	return models.ChattableSlice{changeImage, changeText}
}

func deleteImageCallback(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	messageID := query.Message.MessageID
	fileID := data.MessageFiles[int64(messageID)]
	// delete image from db
	del_req := api.ImageAuthRequest{Image: &api.Image{FileID: fileID}, UserID: userID}
	data.Client.DeleteImage(ctx, &del_req)
	// update message
	return deltaIndexImage(0, ctx, query, data)
}

func randomImageCallback(ctx context.Context, query *tgbotapi.CallbackQuery, data *models.BotData) models.ChattableSlice {
	userID := query.From.ID
	chatID := query.Message.Chat.ID
	messageID := query.Message.MessageID

	image, err := data.Client.GetRandomImage(ctx, &api.Empty{})
	st, ok := status.FromError(err)
	if !ok {
		log.Println(err)
		return models.ChattableSlice{}
	} else if st.Code() == codes.NotFound {
		return models.ChattableSlice{tgbotapi.NewDeleteMessage(chatID, messageID)}
	}
	data.Users[userID].LastDownload = image.FileID
	changeImage := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &keyboards.PublicImageKeyboard,
		},
		Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(image.FileID)),
	}
	caption := utils.PublicImageText(int(image.Upvotes), int(image.Downvotes), image.Description)
	changeText := tgbotapi.NewEditMessageCaption(chatID, messageID, caption)
	changeText.ReplyMarkup = &keyboards.PublicImageKeyboard
	return models.ChattableSlice{changeImage, changeText}
}
