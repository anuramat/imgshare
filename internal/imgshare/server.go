package imgshare

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

type Server struct {
	messageFiles MessageFiles
	api.UnimplementedImgShareServer
	pool   chan struct{}
	DBPool *pgxpool.Pool
}

type MessageFiles struct {
	mu   *sync.RWMutex
	data map[int64]string
}

func NewServer(n_jobs int) (s *Server, err error) {
	s = &Server{}
	s.messageFiles = MessageFiles{&sync.RWMutex{}, make(map[int64]string)}
	s.pool = make(chan struct{}, n_jobs)
	s.DBPool, err = pgxpool.Connect(context.Background(), getDBURL())
	return
}

func getDBURL() (url string) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	url = "postgresql://" + user + ":" + password + "@" + dbhost + ":" + dbport + "/" + dbname
	return url
}
