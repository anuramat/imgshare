package utils

import "strconv"

func PublicImageText(upvotes, downvotes int, description string) string {
	text := "U/D: " + strconv.Itoa(upvotes) + "/" + strconv.Itoa(downvotes) +
		"\nDescription: " + description
	return text
}

func GalleryText(index int32, n_photos, upvotes, downvotes int, description string) string {
	text := "Image " + strconv.Itoa(int(1+index)) + "/" + strconv.Itoa(n_photos) +
		"\nU/D: " + strconv.Itoa(upvotes) + "/" + strconv.Itoa(downvotes) +
		"\nDescription: " + description
	return text
}
