package models

import (
	"strconv"

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

func PublicImageText(upvotes, downvotes int, description string) string {
	text := "U/D: " + strconv.Itoa(upvotes) + "/" + strconv.Itoa(downvotes) + "\nDescription: " + description
	return text
}

func GalleryText(index, n_photos, upvotes, downvotes int, description string) string {
	text := "Image " + strconv.Itoa(1+index) + "/" + strconv.Itoa(n_photos) +
		"\nU/D: " + strconv.Itoa(upvotes) + "/" + strconv.Itoa(downvotes) +
		"\nDescription: " + description
	return text
}

func (d BotData) AddUser(uid int64) (err error) {
	d.Users[uid] = &User{
		State:            StartState,
		LastUpload:       "",
		LastDownload:     "",
		LastGalleryIndex: 0}
	return
}
