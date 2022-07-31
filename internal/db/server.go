package db

import (
	"context"
	"math/rand"
	"sync"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

type Server struct {
	users        Users
	images       Images
	messageFiles MessageFiles
	api.UnimplementedBotDBServer
}

type User struct {
	*api.User
	images []string
}

type Users struct {
	mu   *sync.RWMutex
	data map[int64]*User
}

type Image struct {
	*api.Image
	votes map[int64]bool
}

type Images struct {
	mu   *sync.RWMutex
	data map[string]*Image
}

type MessageFiles struct {
	mu   *sync.RWMutex
	data map[int64]string
}

func NewServer() (s *Server) {
	s = &Server{}
	s.users = Users{&sync.RWMutex{}, make(map[int64]*User)}
	s.images = Images{&sync.RWMutex{}, make(map[string]*Image)}
	s.messageFiles = MessageFiles{&sync.RWMutex{}, make(map[int64]string)}
	return
}

//----- helpers

func (s *Server) getRandomImageID() string {
	images := s.images.data
	k := rand.Intn(len(images))
	for fid := range images {
		if k == 0 {
			return fid
		}
		k--
	}
	panic("unreachable")
}

func (s *Server) upvoteImage(uid int64, fid string) bool {
	images := s.images.data
	if _, ok := images[fid].votes[uid]; ok {
		images[fid].votes[uid] = true
		return true
	} else {
		return false
	}
}

func (s *Server) downvoteImage(uid int64, fid string) bool {
	images := s.images.data
	if _, ok := images[fid].votes[uid]; ok {
		images[fid].votes[uid] = false
		return true
	} else {
		return false
	}
}

func (s *Server) buildImage(fileID string) (result *api.Image) {
	result.FileID = fileID
	result.Description = s.images.data[fileID].Description
	result.Upvotes, result.Downvotes = s.getVotes(fileID)
	return
}

func (s *Server) getVotes(fileID string) (upvotes int64, downvotes int64) {
	votes := s.images.data[fileID].votes
	for _, v := range votes {
		if v {
			upvotes += 1
		} else {
			downvotes += 1
		}
	}
	return
}

//----------- image rpc stuff

func (s *Server) CreateImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.images.mu.Lock()
	defer s.images.mu.Unlock()
	new_image := Image{}
	new_image.Description = input.Image.Description
	new_image.FileID = input.Image.Description
	new_image.votes = make(map[int64]bool)
	s.images.data[input.Image.FileID] = &new_image
	return new_image.Image, nil
}

func (s *Server) ReadImage(_ context.Context, input *api.Image) (*api.Image, error) {
	s.images.mu.RLock()
	defer s.images.mu.RUnlock()
	return s.buildImage(input.FileID), nil
}

func (s *Server) GetRandomImage(_ context.Context, _ *api.Empty) (*api.Image, error) {
	s.images.mu.RLock()
	defer s.images.mu.RUnlock()
	fileID := s.getRandomImageID()
	return s.buildImage(fileID), nil
}

func (s *Server) UpvoteImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.images.mu.Lock()
	defer s.images.mu.Unlock()
	s.upvoteImage(input.UserID, input.Image.FileID)
	return s.buildImage(input.Image.FileID), nil
}

func (s *Server) DownvoteImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.images.mu.Lock()
	defer s.images.mu.Unlock()
	s.downvoteImage(input.UserID, input.Image.FileID)
	return s.buildImage(input.Image.FileID), nil
}

func (s *Server) SetDescriptionImage(_ context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.images.mu.Lock()
	defer s.images.mu.Unlock()
	s.images.data[input.Image.FileID].Description = input.Image.Description
	return s.buildImage(input.Image.FileID), nil
}

func (s *Server) DeleteImage(_ context.Context, input *api.ImageAuthRequest) (_ *api.Empty, err error) {
	s.images.mu.Lock()
	delete(s.images.data, input.Image.FileID)
	s.images.mu.Unlock()

	s.users.mu.Lock()
	user_images := s.users.data[input.UserID].images
	idx := 0
	for i, e := range user_images {
		if e == input.Image.FileID {
			idx = i
		}
	}
	s.users.data[input.UserID].images = append(user_images[:idx], user_images[idx+1:]...)
	s.users.mu.Unlock()
	return
}

//------- user rpc stuff
