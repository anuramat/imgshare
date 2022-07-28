package commands

import (
	"errors"

	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func CommandRouter(cmd string, userID int64, users models.Users) (err error) {
	// switches to a respective state
	switch cmd {
	case "start":
		users[userID].State = models.StartState
		return
	case "upload":
		users[userID].State = models.UploadImageInitState
		return
	case "gallery":
		users[userID].State = models.GalleryState
		return
	case "random":
		users[userID].State = models.RandomImageState
		return
	}
	err = errors.New("invalid command")
	return
}
