package db

import (
	"context"
	"errors"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

func (s *Server) CreateImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.Lock()
	defer s.images.mu.Unlock()

	new_image := NewImage()
	new_image.Description = input.Image.Description
	new_image.FileID = input.Image.FileID
	s.images.data[input.Image.FileID] = &new_image
	return new_image.Image, nil
}

func (s *Server) ReadImage(_ context.Context, input *api.Image) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.RLock()
	defer s.images.mu.RUnlock()

	return s.buildImage(input.FileID)
}

func (s *Server) GetRandomImage(_ context.Context, _ *api.Empty) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.RLock()
	defer s.images.mu.RUnlock()

	fileID, err := s.getRandomImageID()
	if err != nil {
		return nil, err
	}
	return s.buildImage(fileID)
}

func (s *Server) UpvoteImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.Lock()
	defer s.images.mu.Unlock()

	err := s.upvoteImage(input.UserID, input.Image.FileID)
	if err != nil {
		return nil, err
	}
	return s.buildImage(input.Image.FileID)
}

func (s *Server) DownvoteImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.Lock()
	defer s.images.mu.Unlock()

	err := s.downvoteImage(input.UserID, input.Image.FileID)
	if err != nil {
		return nil, err
	}
	return s.buildImage(input.Image.FileID)
}

func (s *Server) SetDescriptionImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.Lock()
	defer s.images.mu.Unlock()

	if _, ok := s.images.data[input.Image.FileID]; !ok {
		return nil, errors.New("image not found")
	}
	s.images.data[input.Image.FileID].Description = input.Image.Description
	return s.buildImage(input.Image.FileID)
}

func (s *Server) DeleteImage(_ context.Context, input *api.ImageAuthRequest) (_ *api.Empty, err error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.Lock()
	delete(s.images.data, input.Image.FileID)
	s.images.mu.Unlock()

	s.users.mu.Lock()
	defer s.users.mu.Unlock()
	user_images := s.users.data[input.UserID].images
	idx := -1
	for i, e := range user_images {
		if e == input.Image.FileID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, errors.New("image not found in users gallery")
	}
	s.users.data[input.UserID].images = append(user_images[:idx], user_images[idx+1:]...)
	return
}

func (s *Server) GetAllImages(_ context.Context, _ *api.Empty) (*api.Images, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	s.images.mu.Lock()
	defer s.images.mu.Unlock()

	images_slice := make([]*api.Image, len(s.images.data))
	i := 0
	for _, v := range s.images.data {
		images_slice[i] = v.Image
		i += 1
	}
	return &api.Images{Image: images_slice}, nil
}
