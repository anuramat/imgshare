package imgshare

import (
	"context"
	"errors"
	"log"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
)

func (s *Server) CreateImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Empty, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	sql := "INSERT INTO images (fileid, userid, description) VALUES ($1, $2, $3);"
	s.DBPool.Exec(ctx, sql, input.Image.FileID, input.UserID, input.Image.Description)
	return &api.Empty{}, nil
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
	s.DBPool.Exec(ctx, sql, fid, uid, vote)
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
	s.DBPool.Exec(ctx, sql, input.Image.Description, input.Image.FileID, input.UserID)
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
	s.DBPool.Exec(ctx, sql_images, input.Image.FileID, input.UserID)
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

	sql := `SELECT
	images.fileid as fileid,
	images.description as description,
	COUNT(*) FILTER (WHERE votes.upvote) as upvotes,
	COUNT(*) FILTER (WHERE NOT votes.upvote) as downvotes
	FROM images
	LEFT JOIN votes ON images.fileid = votes.fileid
	GROUP BY images.fileid
	ORDER BY images.fileid
	LIMIT $1
	OFFSET $2;`
	rows, err := s.DBPool.Query(ctx, sql, page.Limit, page.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (s *Server) buildImage(ctx context.Context, fileID string) (*api.Image, error) {
	result := &api.Image{}

	sql := `SELECT
	images.fileid as fileid,
	images.description as description,
	COUNT(*) FILTER (WHERE votes.upvote) as upvotes,
	COUNT(*) FILTER (WHERE NOT votes.upvote) as downvotes
	FROM images
	LEFT JOIN votes ON images.fileid = votes.fileid
	WHERE images.fileid = $1
	GROUP BY images.fileid`

	row := s.DBPool.QueryRow(ctx, sql, fileID)
	err := row.Scan(&result.FileID, &result.Description, &result.Upvotes, &result.Downvotes)

	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *Server) GetGalleryImage(ctx context.Context, request *api.GalleryRequest) (*api.GalleryImage, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	default:
	}

	var total int32 = 0
	// get number of images in the gallry
	sql_count := `SELECT
	COUNT(*) AS count
	FROM images
	WHERE userid = $1`
	row_count := s.DBPool.QueryRow(ctx, sql_count, request.UserID)
	err := row_count.Scan(&total)
	if err != nil {
		return nil, err
	}
	offset := request.Offset
	if offset >= total {
		offset = total - 1
	}

	image := &api.Image{}

	sql_get := `SELECT
	images.fileid as fileid,
	images.description as description,
	COUNT(*) FILTER (WHERE votes.upvote) as upvotes,
	COUNT(*) FILTER (WHERE NOT votes.upvote) as downvotes
	FROM images
	LEFT JOIN votes ON images.fileid = votes.fileid
	WHERE images.userid = $1
	GROUP BY images.fileid
	ORDER BY images.fileid
	LIMIT 1
	OFFSET $2`

	row_get := s.DBPool.QueryRow(ctx, sql_get, request.UserID, offset)
	err = row_get.Scan(&image.FileID, &image.Description, &image.Upvotes, &image.Downvotes)

	if err != nil {
		return nil, err
	}
	result := &api.GalleryImage{}
	result.Image = image
	result.Offset = offset
	result.Total = total
	return result, nil
}
