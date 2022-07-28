package models

import (
	"math/rand"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// message id -> file id
type MessageFiles map[int64]string
type Users map[int64]*User
type Images map[string]*ImageMeta
type ChattableSlice []tgbotapi.Chattable
type state uint8

type User struct {
	State            state
	LastUpload       string
	LastDownload     string
	Images           []string
	LastGalleryIndex int
}

// HACK votes are gonna be very slow very soon
// TODO rewrite before it's too late
type ImageMeta struct {
	Description string
	Votes       map[int64]bool
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

// TODO rewrite
func GetRandomImage(images Images) string {
	k := rand.Intn(len(images))
	for fid := range images {
		if k == 0 {
			return fid
		}
		k--
	}
	panic("unreachable")
}

func GetVotes(votes map[int64]bool) (upvotes int, downvotes int) {
	for _, v := range votes {
		if v {
			upvotes += 1
		} else {
			downvotes += 1
		}
	}
	return
}

func UpvoteImage(images Images, uid int64, fid string) bool {
	if !images[fid].Votes[uid] {
		images[fid].Votes[uid] = true
		return true
	} else {
		return false
	}
}

func DownvoteImage(images Images, uid int64, fid string) bool {
	if images[fid].Votes[uid] {
		images[fid].Votes[uid] = false
		return true
	} else {
		return false
	}
}

func PublicImageText(fileID string, images Images) string {
	meta := images[fileID]
	upvotes, downvotes := GetVotes(meta.Votes)
	text := "U/D: " + strconv.Itoa(upvotes) + "/" + strconv.Itoa(downvotes) + "\nDescription: " + meta.Description
	return text
}

func GalleryText(index int, fileID string, userID int64, images Images, users Users) string {
	u := users[userID]
	n_photos := len(u.Images)
	meta := images[fileID]
	upvotes, downvotes := GetVotes(meta.Votes)
	text := "Image " + strconv.Itoa(1+index) + "/" + strconv.Itoa(n_photos) +
		"\nU/D: " + strconv.Itoa(upvotes) + "/" + strconv.Itoa(downvotes) +
		"\nDescription: " + meta.Description
	return text
}
