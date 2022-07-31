package db

import (
	"sync"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

type Server struct {
	users        Users
	images       Images
	messageFiles MessageFiles
	api.UnimplementedBotDBServer
	pool chan struct{}
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

func NewServer(n_jobs int) (s *Server) {
	s = &Server{}
	s.users = Users{&sync.RWMutex{}, make(map[int64]*User)}
	s.images = Images{&sync.RWMutex{}, make(map[string]*Image)}
	s.messageFiles = MessageFiles{&sync.RWMutex{}, make(map[int64]string)}
	s.pool = make(chan struct{}, n_jobs)
	return
}

func NewImage() Image {
	return Image{&api.Image{}, make(map[int64]bool)}
}

func NewUser() User {
	return User{&api.User{}, make([]string, 3)}
}
