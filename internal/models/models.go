package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

// message id -> file id
type MessageFiles map[int64]string
type Users map[int64]*User
type ChattableSlice []tgbotapi.Chattable
type state uint8

type BotData struct {
	Client api.ImgShareClient
	Users
	MessageFiles
}

type User struct {
	State            state
	LastUpload       string
	LastDownload     string
	LastGalleryIndex int
}

const (
	StartState state = iota
	NoState
	UploadImageInitState
	UploadImageState
	UploadDescriptionState
	EditDescriptionState
	RandomImageState
	GalleryState
)

func NewBotData(client api.ImgShareClient) *BotData {
	return &BotData{
		Client:       client,
		Users:        Users{},
		MessageFiles: MessageFiles{},
	}
}

func (d BotData) AddUser(uid int64) (err error) {
	d.Users[uid] = &User{
		State:            StartState,
		LastUpload:       "",
		LastDownload:     "",
		LastGalleryIndex: 0}
	return
}
