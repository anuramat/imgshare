package imgshare

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

type Server struct {
	api.UnimplementedImgShareServer
	pool   chan struct{}
	dbpool dbpool
}

type dbpool interface {
	Close()
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func NewServer(n_jobs int) (s *Server, err error) {
	s = &Server{}
	s.pool = make(chan struct{}, n_jobs)
	s.dbpool, err = pgxpool.Connect(context.Background(), getDBURL())
	return
}

func (s Server) Close() {
	s.dbpool.Close()
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
