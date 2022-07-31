package db

import (
	"errors"
	"math/rand"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

func (s *Server) getRandomImageID() (string, error) {
	images := s.images.data
	if len(images) < 1 {
		return "", errors.New("no images yet in DB")
	}
	k := rand.Intn(len(images))
	for fid := range images {
		if k == 0 {
			return fid, nil
		}
		k--
	}
	panic("unreachable")
}

func (s *Server) upvoteImage(uid int64, fid string) error {
	images := s.images.data
	if _, ok := images[fid]; ok {
		images[fid].votes[uid] = true
		return nil
	} else {
		return errors.New("image not found")
	}
}

func (s *Server) downvoteImage(uid int64, fid string) error {
	images := s.images.data
	if _, ok := images[fid]; ok {
		images[fid].votes[uid] = false
		return nil
	} else {
		return errors.New("image not found")
	}
}

func (s *Server) buildImage(fileID string) (result *api.Image, err error) {
	result = &api.Image{}
	result.FileID = fileID
	if _, ok := s.images.data[fileID]; !ok {
		return nil, errors.New("fileID wasn't found or properly created/updated")
	}
	result.Description = s.images.data[fileID].Description
	result.Upvotes, result.Downvotes, err = s.getVotes(fileID)
	return
}

func (s *Server) getVotes(fileID string) (upvotes int64, downvotes int64, err error) {
	if _, ok := s.images.data[fileID]; !ok {
		return 0, 0, errors.New("image not found")
	}
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
