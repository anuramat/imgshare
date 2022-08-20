package apierr

import "errors"

var (
	ErrNoImages = errors.New("no images found")
	ErrTimeout  = errors.New("timeout")
)
