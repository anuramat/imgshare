package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var GalleryKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("previous", PreviousImageButton),
		tgbotapi.NewInlineKeyboardButtonData("next", NextImageButton),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("change description", EditDescriptionButton),
		tgbotapi.NewInlineKeyboardButtonData("delete", DeleteImageButton),
	),
)

var PublicImageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("feeling lucky", RandomImageButton),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("upvote", UpvoteImageButton),
		tgbotapi.NewInlineKeyboardButtonData("downvote", DownvoteImageButton),
	),
)

const (
	UpvoteImageButton     = "upvote"
	DownvoteImageButton   = "downvote"
	NextImageButton       = "nextImage"
	PreviousImageButton   = "previousImage"
	RandomImageButton     = "randomImage"
	EditDescriptionButton = "editDescription"
	DeleteImageButton     = "deleteImage"
)
