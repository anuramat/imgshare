package imgshare

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/apierr"
	"gitlab.ozon.dev/anuramat/homework-1/internal/mocks"
)

var (
	Err = errors.New("")
)

func TestServer_CreateImage(t *testing.T) {
	t.Run("good input", func(t *testing.T) {
		// arrange
		var uid int64 = 124123
		fid := "1452asdfad123"
		description := "dsafasdfasdf"
		var upvotes int64 = 123
		var downvotes int64 = 124123

		s := &Server{}
		s.pool = make(chan struct{}, 1)
		mockCtrl := gomock.NewController(t)
		dbpool := mocks.NewMockDBPool(mockCtrl)
		ctx := context.Background()
		input := &api.ImageAuthRequest{UserID: uid, Image: &api.Image{FileID: fid, Description: description, Upvotes: upvotes, Downvotes: downvotes}}
		dbpool.EXPECT().Exec(ctx, "INSERT INTO images (fileid, userid, description) VALUES ($1, $2, $3);", fid, uid, description).Return(nil, nil)
		s.dbpool = dbpool

		// act
		_, err := s.CreateImage(ctx, input)

		// assert
		assert.Nil(t, err)
		assert.Empty(t, s.pool)
	})

	t.Run("db error case", func(t *testing.T) {
		// arrange
		var uid int64 = 124123
		fid := "1452asdfad123"
		description := "dsafasdfasdf"
		var upvotes int64 = 123
		var downvotes int64 = 124123

		s := &Server{}
		s.pool = make(chan struct{}, 1)
		mockCtrl := gomock.NewController(t)
		dbpool := mocks.NewMockDBPool(mockCtrl)
		ctx := context.Background()
		input := &api.ImageAuthRequest{UserID: uid, Image: &api.Image{FileID: fid, Description: description, Upvotes: upvotes, Downvotes: downvotes}}
		dbpool.EXPECT().Exec(ctx, "INSERT INTO images (fileid, userid, description) VALUES ($1, $2, $3);", fid, uid, description).Return(nil, Err)
		s.dbpool = dbpool

		// act
		_, err := s.CreateImage(ctx, input)

		// assert
		assert.NotNil(t, err)
		assert.Empty(t, s.pool)
	})

	t.Run("timeout case", func(t *testing.T) {
		// arrange
		var uid int64 = 124123
		fid := "1452asdfad123"
		description := "dsafasdfasdf"
		var upvotes int64 = 123
		var downvotes int64 = 124123

		s := &Server{}
		s.pool = make(chan struct{}, 1)
		mockCtrl := gomock.NewController(t)
		dbpool := mocks.NewMockDBPool(mockCtrl)
		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		cancel()
		input := &api.ImageAuthRequest{UserID: uid, Image: &api.Image{FileID: fid, Description: description, Upvotes: upvotes, Downvotes: downvotes}}
		s.dbpool = dbpool

		// act
		_, err := s.CreateImage(ctx, input)

		// assert
		assert.Equal(t, apierr.ErrTimeout, err)
		assert.Empty(t, s.pool)
	})
}

func TestServer_ReadImage(t *testing.T) {
}

func TestServer_GetRandomImage(t *testing.T) {
}

func TestServer_upsertVoteImage(t *testing.T) {
}

func TestServer_UpvoteImage(t *testing.T) {
}

func TestServer_DownvoteImage(t *testing.T) {
}

func TestServer_SetDescriptionImage(t *testing.T) {
}

func TestServer_DeleteImage(t *testing.T) {
}

func TestServer_GetAllImages(t *testing.T) {
}

func TestServer_buildImage(t *testing.T) {
}

func TestServer_GetGalleryImage(t *testing.T) {
}
