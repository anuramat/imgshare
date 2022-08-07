package imgshare

import (
	"context"
	"errors"
	"log"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

func (s *Server) CreateImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	sql := "INSERT INTO images (fileid, userid, description) VALUES ($1, $2, $3);"
	s.dbPool.QueryRow(ctx, sql, input.Image.FileID, input.UserID, input.Image.Description)
	return s.buildImage(ctx, input.Image.FileID)
}

func (s *Server) ReadImage(ctx context.Context, input *api.Image) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	return s.buildImage(ctx, input.FileID)
}

func (s *Server) GetRandomImage(ctx context.Context, _ *api.Empty) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	// TODO get random fileID, put in fileID
	// return s.buildImage(ctx, fileID)
	return nil, nil
}

func (s *Server) upsertVoteImage(ctx context.Context, uid int64, fid string, vote bool) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	sql := "INSERT INTO votes (fileid, userid, upvote) VALUES ($1, $2, $3) ON CONFLICT (fileid, userid) DO UPDATE SET upvote = $3;"
	s.dbPool.QueryRow(ctx, sql, fid, uid, vote)
	return s.buildImage(ctx, fid)
}

func (s *Server) UpvoteImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	return s.upsertVoteImage(ctx, input.UserID, input.Image.FileID, true)
}

func (s *Server) DownvoteImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	return s.upsertVoteImage(ctx, input.UserID, input.Image.FileID, false)
}

func (s *Server) SetDescriptionImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	sql := "UPDATE images SET description = $1 WHERE fileid = $2 AND userid = $3;" // TODO check if updated
	s.dbPool.QueryRow(ctx, sql, input.Image.Description, input.Image.FileID, input.UserID)
	return s.buildImage(ctx, input.Image.FileID) // TODO return rows instead of building
}

func (s *Server) DeleteImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Empty, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	sql_images := "DELETE FROM images WHERE fileid = $1 AND userid = $2;"
	s.dbPool.QueryRow(ctx, sql_images, input.Image.FileID, input.UserID)
	// TODO delete votes? only if the image itself was deleted
	return &api.Empty{}, nil
}

func (s *Server) GetAllImages(ctx context.Context, page *api.Page) (*api.Images, error) {
	// HW-2 requirement
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	sql := `SELECT images.fileid, images.description,
	COUNT(*) FILTER (WHERE votes.upvote) as upvotes,
	COUNT(*) FILTER (WHERE NOT votes.upvote) as downvotes
	FROM images
	LEFT JOIN votes ON images.fileid = votes.fileid
	GROUP BY images.fileid
	ORDER BY images.fileid
	LIMIT $1
	OFFSET $2;`
	rows, err := s.dbPool.Query(ctx, sql, page.Limit, page.Offset)
	if err != nil {
		return nil, err
	}

	images_slice := make([]*api.Image, 0, page.Limit)
	var img *api.Image
	for rows.Next() {
		img = &api.Image{}
		rows.Scan(&img.FileID, &img.Description, &img.Upvotes, &img.Downvotes)
		images_slice = append(images_slice, img)
	}
	log.Println(images_slice)
	return &api.Images{Image: images_slice}, nil
}

func (s *Server) buildImage(ctx context.Context, fileID string) (result *api.Image, err error) {
	result = &api.Image{}

	sql := "SELECT fileid, description FROM images WHERE fileid = $1;"
	row := s.dbPool.QueryRow(ctx, sql, fileID)
	row.Scan(&result.FileID, &result.Description)

	sql_up := "SELECT COUNT(*) FROM votes WHERE fileid = $1 AND upvote = TRUE;"
	row_up := s.dbPool.QueryRow(ctx, sql_up, fileID)
	row_up.Scan(&result.Upvotes)

	sql_down := "SELECT COUNT(*) FROM votes WHERE fileid = $1 and upvote = FALSE;"
	row_down := s.dbPool.QueryRow(ctx, sql_down, fileID)
	row_down.Scan(&result.Downvotes)

	return
}
