package messages

import (
	"context"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/mocks"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func makeMsg(uid int64, cid int64, fid string, text string) *tgbotapi.Message {
	return &tgbotapi.Message{
		Photo: []tgbotapi.PhotoSize{{FileID: fid}},
		From:  &tgbotapi.User{ID: uid},
		Chat:  &tgbotapi.Chat{ID: cid},
		Text:  text,
	}
}

func TestUploadImageHandler(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 123
		var cid int64 = 456
		var fid string = "789"

		msg := makeMsg(uid, cid, fid, "")
		ctx := context.Background()
		data := models.NewBotData(nil)
		data.AddUser(uid)

		// act
		_ = UploadImageHandler(ctx, msg, data)

		// assert
		assert.Equal(t, data.Users[uid].LastUpload, fid)
		assert.Equal(t, data.Users[uid].State, models.UploadDescriptionState)
	})
}

func TestUploadDescriptionHandler(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 123
		var cid int64 = 456
		var fid string = "789"
		var text string = "xdfasdf"

		msg := makeMsg(uid, cid, "", text)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.ImageAuthRequest{UserID: msg.From.ID, Image: &api.Image{FileID: fid, Description: msg.Text}}
		mockClient.EXPECT().CreateImage(ctx, gomock.Eq(&req)).Return(nil, nil)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.Users[msg.From.ID].LastUpload = fid

		// act
		_ = UploadDescriptionHandler(ctx, msg, data)

		// assert
		assert.Equal(t, data.Users[uid].State, models.NoState)
	})
}

func TestEditDescriptionHandler(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 123
		var cid int64 = 456
		var fid string = "789"
		var text string = "xdfasdf"

		msg := makeMsg(uid, cid, "", text)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.ImageAuthRequest{UserID: msg.From.ID, Image: &api.Image{FileID: fid, Description: msg.Text}}
		mockClient.EXPECT().SetDescriptionImage(ctx, gomock.Eq(&req)).Return(nil, nil)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.Users[msg.From.ID].LastUpload = fid

		// act
		_ = EditDescriptionHandler(ctx, msg, data)

		// assert
		assert.Equal(t, data.Users[uid].State, models.NoState)
	})
}

func TestRandomImageHandler(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 123
		var cid int64 = 456
		var fid string = "789"

		msg := makeMsg(uid, cid, "", "")
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.Empty{}
		resp := &api.Image{FileID: fid}
		mockClient.EXPECT().GetRandomImage(ctx, gomock.Eq(&req)).Return(resp, nil)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)

		// act
		photo := RandomImageHandler(ctx, msg, data)

		// assert
		assert.Equal(t, data.Users[uid].LastDownload, fid)
		assert.IsType(t, tgbotapi.PhotoConfig{}, photo)
	})
}

func TestGalleryHandler(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 123
		var cid int64 = 456
		var fid string = "789"
		var last_index int = 1453
		var offset int = 5123
		var total int = 124123

		msg := makeMsg(uid, cid, "", "")
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.GalleryRequest{Offset: int32(last_index), UserID: uid}
		resp := &api.GalleryImage{Offset: int32(offset), Total: int32(total), Image: &api.Image{FileID: fid}}
		mockClient.EXPECT().GetGalleryImage(ctx, gomock.Eq(&req)).Return(resp, nil)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.Users[uid].LastGalleryIndex = last_index

		// act
		photo := GalleryHandler(ctx, msg, data)

		// assert
		assert.Equal(t, data.Users[uid].LastGalleryIndex, offset)
		assert.Equal(t, data.Users[uid].LastDownload, fid)
		assert.IsType(t, tgbotapi.PhotoConfig{}, photo)
	})
}
