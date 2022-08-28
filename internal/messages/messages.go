package messages

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

func DefaultHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(msg.Chat.ID, "Choose a command from the menu.")
}

func StartHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.MessageConfig {
	data.Users[msg.From.ID].State = models.NoState
	return tgbotapi.NewMessage(msg.Chat.ID, "Welcome! I guess it would be nice if I told you what this bot does, but I'm too sleep deprived.")
}

func UploadImageInitHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.MessageConfig {
	data.Users[msg.From.ID].State = models.UploadImageState
	return tgbotapi.NewMessage(msg.Chat.ID, "Send me an image!")
}

func UploadImageHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.MessageConfig {
	if len(msg.Photo) == 0 {
		return tgbotapi.NewMessage(msg.Chat.ID, "Invalid image, try again.")
	}
	fid := string(msg.Photo[len(msg.Photo)-1].FileID)
	data.Users[msg.From.ID].LastUpload = fid
	data.Users[msg.From.ID].State = models.UploadDescriptionState
	return tgbotapi.NewMessage(msg.Chat.ID, "Send the description for your image")
}

func UploadDescriptionHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.MessageConfig {
	fid := data.Users[msg.From.ID].LastUpload
	req := api.ImageAuthRequest{UserID: msg.From.ID, Image: &api.Image{FileID: fid, Description: msg.Text}}
	_, err := data.Client.CreateImage(ctx, &req)
	if err != nil {
		log.Println(err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Error!")
	}

	data.Users[msg.From.ID].State = models.NoState
	return tgbotapi.NewMessage(msg.Chat.ID, "Uploaded successfully!")
}

func EditDescriptionHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.MessageConfig {
	fid := data.Users[msg.From.ID].LastUpload
	req := api.ImageAuthRequest{UserID: msg.From.ID, Image: &api.Image{FileID: fid, Description: msg.Text}}
	_, err := data.Client.SetDescriptionImage(ctx, &req)
	if err != nil {
		log.Println(err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Error!")
	}

	data.Users[msg.From.ID].State = models.NoState
	return tgbotapi.NewMessage(msg.Chat.ID, "Description changed successfully!")
}

func RandomImageHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.Chattable {
	result, err := data.Client.GetRandomImage(ctx, &api.Empty{})

	st, ok := status.FromError(err)
	if !ok {
		log.Println(err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Error!")
	} else if st.Code() == codes.NotFound {
		return tgbotapi.NewMessage(msg.Chat.ID, "No images yet!")
	}

	text := utils.PublicImageText(int(result.Upvotes), int(result.Downvotes), result.Description)
	photo := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileID(result.FileID))
	photo.Caption = text
	photo.ReplyMarkup = keyboards.PublicImageKeyboard
	data.Users[msg.From.ID].LastDownload = result.FileID
	return photo
}

func GalleryHandler(ctx context.Context, msg *tgbotapi.Message, data *models.BotData) tgbotapi.Chattable {
	index := data.Users[msg.From.ID].LastGalleryIndex
	req := api.GalleryRequest{Offset: int32(index), UserID: msg.From.ID}
	result, err := data.Client.GetGalleryImage(ctx, &req)
	st, ok := status.FromError(err)
	if !ok {
		log.Println(err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Error!")
	} else if st.Code() == codes.NotFound {
		return tgbotapi.NewMessage(msg.Chat.ID, "No images yet!")
	}
	text := utils.GalleryText(result.Offset, int(result.Total), int(result.Image.Upvotes), int(result.Image.Downvotes), result.Image.Description)
	photo := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileID(result.Image.FileID))
	photo.Caption = text
	photo.ReplyMarkup = keyboards.GalleryKeyboard
	data.Users[msg.From.ID].LastDownload = result.Image.FileID
	data.Users[msg.From.ID].LastGalleryIndex = result.Offset
	return photo
}
