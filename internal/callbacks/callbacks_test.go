package callbacks

import (
	"context"
	"errors"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/mocks"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func makeQuery(uid, cid int64, mid int) *tgbotapi.CallbackQuery {
	return &tgbotapi.CallbackQuery{
		From: &tgbotapi.User{ID: uid},
		Message: &tgbotapi.Message{
			MessageID: mid,
			Chat:      &tgbotapi.Chat{ID: cid},
		},
	}
}

var (
	Err = errors.New("")
)

func Test_upvoteCallback(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var upvotes int64 = 123413
		var downvotes int64 = 19495
		var mid int = 124123
		var fid string = "asdf"
		var description string = "adsfhguhe4 4rog"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.ImageAuthRequest{Image: &api.Image{FileID: fid}, UserID: uid}
		resp := api.Image{FileID: fid, Upvotes: upvotes, Downvotes: downvotes, Description: description}
		mockClient.EXPECT().UpvoteImage(ctx, gomock.Eq(&req)).Return(&resp, nil)
		data := models.NewBotData(mockClient)
		data.MessageFiles[mid] = fid

		// act
		_ = upvoteCallback(ctx, query, data)
	})

	t.Run("grpc error case", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var mid int = 124123
		var fid string = "asdf"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.ImageAuthRequest{Image: &api.Image{FileID: fid}, UserID: uid}
		mockClient.EXPECT().UpvoteImage(ctx, gomock.Eq(&req)).Return(nil, Err)
		data := models.NewBotData(mockClient)
		data.MessageFiles[mid] = fid

		// act
		_ = upvoteCallback(ctx, query, data)
	})
}

func Test_downvoteCallback(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var upvotes int64 = 123413
		var downvotes int64 = 19495
		var mid int = 124123
		var fid string = "asdf"
		var description string = "adsfhguhe4 4rog"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.ImageAuthRequest{Image: &api.Image{FileID: fid}, UserID: uid}
		resp := api.Image{FileID: fid, Upvotes: upvotes, Downvotes: downvotes, Description: description}
		mockClient.EXPECT().DownvoteImage(ctx, gomock.Eq(&req)).Return(&resp, nil)
		data := models.NewBotData(mockClient)
		data.MessageFiles[mid] = fid

		// act
		_ = downvoteCallback(ctx, query, data)
	})

	t.Run("grpc error case", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var mid int = 124123
		var fid string = "asdf"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.ImageAuthRequest{Image: &api.Image{FileID: fid}, UserID: uid}
		mockClient.EXPECT().DownvoteImage(ctx, gomock.Eq(&req)).Return(nil, Err)
		data := models.NewBotData(mockClient)
		data.MessageFiles[mid] = fid

		// act
		_ = downvoteCallback(ctx, query, data)
	})
}

func Test_editDescriptionCallback(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var mid int = 124123
		var fid string = "asdf"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		data := models.NewBotData(nil)
		data.AddUser(uid)
		data.MessageFiles[mid] = fid

		// act
		_ = editDescriptionCallback(ctx, query, data)

		// assert
		assert.Equal(t, fid, data.Users[uid].LastUpload)
		assert.Equal(t, models.EditDescriptionState, data.Users[uid].State)
	})
}

func Test_deltaIndexImage(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var upvotes int64 = 123413
		var downvotes int64 = 19495
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"
		var description string = "adsfhguhe4 4rog"
		var total int32 = 12312312
		var delta int32 = 1

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.GalleryRequest{Offset: index + delta, UserID: uid}
		resp := api.GalleryImage{Offset: index + delta, Total: total, Image: &api.Image{FileID: fid, Upvotes: upvotes, Downvotes: downvotes, Description: description}}
		mockClient.EXPECT().GetGalleryImage(ctx, gomock.Eq(&req)).Return(&resp, nil)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.Users[uid].LastGalleryIndex = index
		data.MessageFiles[mid] = fid

		// act
		_ = deltaIndexImage(delta, ctx, query, data)

		// assert
		assert.Equal(t, data.Users[uid].LastGalleryIndex, index+delta)
		assert.Equal(t, data.Users[uid].LastDownload, fid)
	})

	t.Run("negative index", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var upvotes int64 = 123413
		var downvotes int64 = 19495
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"
		var description string = "adsfhguhe4 4rog"
		var total int32 = 12312312
		var delta int32 = -index - 100

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.GalleryRequest{Offset: 0, UserID: uid}
		resp := api.GalleryImage{Offset: 0, Total: total, Image: &api.Image{FileID: fid, Upvotes: upvotes, Downvotes: downvotes, Description: description}}
		mockClient.EXPECT().GetGalleryImage(ctx, gomock.Eq(&req)).Return(&resp, nil)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.Users[uid].LastGalleryIndex = index
		data.MessageFiles[mid] = fid

		// act
		_ = deltaIndexImage(delta, ctx, query, data)

		// assert
		assert.Equal(t, data.Users[uid].LastGalleryIndex, int32(0))
		assert.Equal(t, data.Users[uid].LastDownload, fid)
	})

	t.Run("grpc error case", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"
		var delta int32 = 1

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)
		req := api.GalleryRequest{Offset: index + delta, UserID: uid}
		mockClient.EXPECT().GetGalleryImage(ctx, gomock.Eq(&req)).Return(nil, Err)
		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.Users[uid].LastGalleryIndex = index
		data.MessageFiles[mid] = fid

		// act
		_ = deltaIndexImage(delta, ctx, query, data)
	})
}

func Test_deleteImageCallback(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var upvotes int64 = 123413
		var downvotes int64 = 19495
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"
		var description string = "adsfhguhe4 4rog"
		var total int32 = 12312312
		var delta int32 = 0

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)

		req := api.ImageAuthRequest{Image: &api.Image{FileID: fid}, UserID: uid}
		mockClient.EXPECT().DeleteImage(ctx, gomock.Eq(&req)).Return(nil, nil)

		req_gall := api.GalleryRequest{Offset: index + delta, UserID: uid}
		resp_gall := api.GalleryImage{Offset: index + delta, Total: total, Image: &api.Image{FileID: fid, Upvotes: upvotes, Downvotes: downvotes, Description: description}}
		mockClient.EXPECT().GetGalleryImage(ctx, gomock.Eq(&req_gall)).Return(&resp_gall, nil)

		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.MessageFiles[mid] = fid
		data.Users[uid].LastGalleryIndex = index

		// act
		_ = deleteImageCallback(ctx, query, data)
	})

	t.Run("grpc error case", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)

		req := api.ImageAuthRequest{Image: &api.Image{FileID: fid}, UserID: uid}
		mockClient.EXPECT().DeleteImage(ctx, gomock.Eq(&req)).Return(nil, Err)

		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.MessageFiles[mid] = fid
		data.Users[uid].LastGalleryIndex = index

		// act
		_ = deleteImageCallback(ctx, query, data)
	})
}

func Test_randomImageCallback(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var upvotes int64 = 123413
		var downvotes int64 = 19495
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"
		var description string = "adsfhguhe4 4rog"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)

		resp := api.Image{FileID: fid, Upvotes: upvotes, Downvotes: downvotes, Description: description}
		mockClient.EXPECT().GetRandomImage(ctx, gomock.Eq(&api.Empty{})).Return(&resp, nil)

		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.MessageFiles[mid] = fid
		data.Users[uid].LastGalleryIndex = index

		// act
		_ = randomImageCallback(ctx, query, data)
	})

	t.Run("grpc error case", func(t *testing.T) {
		// arrange
		var uid int64 = 1233452
		var cid int64 = 1251234
		var index int32 = 1123
		var mid int = 124123
		var fid string = "asdf"

		query := makeQuery(uid, cid, mid)
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		mockClient := mocks.NewMockImgShareClient(mockCtrl)

		mockClient.EXPECT().GetRandomImage(ctx, gomock.Eq(&api.Empty{})).Return(nil, Err)

		data := models.NewBotData(mockClient)
		data.AddUser(uid)
		data.MessageFiles[mid] = fid
		data.Users[uid].LastGalleryIndex = index

		// act
		_ = randomImageCallback(ctx, query, data)
	})
}
