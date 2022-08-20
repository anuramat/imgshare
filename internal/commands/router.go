package commands

import (
	"errors"

	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func CommandRouter(cmd string, userID int64, data *models.BotData) (err error) {
	// switches to a respective state
	switch cmd {
	case "start":
		data.Users[userID].State = models.StartState
		return
	case "upload":
		data.Users[userID].State = models.UploadImageInitState
		return
	case "gallery":
		data.Users[userID].State = models.GalleryState
		return
	case "random":
		data.Users[userID].State = models.RandomImageState
		return
	}
	err = errors.New("invalid command")
	return
}
