package imgshare

import (
	"context"
	"log"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/apierr"
)

func (s *Server) CreateImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Empty, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	sql := "INSERT INTO images (fileid, userid, description) VALUES ($1, $2, $3);"
	_, err := s.DBPool.Exec(ctx, sql, input.Image.FileID, input.UserID, input.Image.Description)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

func (s *Server) ReadImage(ctx context.Context, input *api.Image) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	return s.buildImage(ctx, input.FileID)
}

func (s *Server) GetRandomImage(ctx context.Context, _ *api.Empty) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	var fileID string
	var count int

	sql_count := "SELECT COUNT(*) FROM images;"
	row_count := s.DBPool.QueryRow(ctx, sql_count)
	err_count := row_count.Scan(&count)
	if err_count != nil {
		return nil, err_count
	}

	if count == 0 {
		return nil, apierr.ErrNoImages
	}

	sql_random := "SELECT fileid FROM images ORDER BY random() LIMIT 1;"
	row_random := s.DBPool.QueryRow(ctx, sql_random)
	err_random := row_random.Scan(&fileID)
	if err_random != nil {
		return nil, err_random
	}

	return s.buildImage(ctx, fileID)
}

func (s *Server) upsertVoteImage(ctx context.Context, uid int64, fid string, vote bool) (*api.Image, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	sql := "INSERT INTO votes (fileid, userid, upvote) VALUES ($1, $2, $3) ON CONFLICT (fileid, userid) DO UPDATE SET upvote = $3;"
	_, err := s.DBPool.Exec(ctx, sql, fid, uid, vote)
	if err != nil {
		return nil, err
	}
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
		return nil, apierr.ErrTimeout
	default:
	}

	sql := "UPDATE images SET description = $1 WHERE fileid = $2 AND userid = $3;"
	_, err := s.DBPool.Exec(ctx, sql, input.Image.Description, input.Image.FileID, input.UserID)
	if err != nil {
		return nil, err
	}

	return s.buildImage(ctx, input.Image.FileID)
}

func (s *Server) DeleteImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Empty, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	sql_images := "DELETE FROM images WHERE fileid = $1 AND userid = $2;"
	res, err_images := s.DBPool.Exec(ctx, sql_images, input.Image.FileID, input.UserID)
	if err_images != nil {
		return nil, err_images
	}
	count := res.RowsAffected()
	if count == 0 {
		return &api.Empty{}, nil
	}

	sql_votes := "DELETE FROM votes WHERE fileid = $1;"
	_, err_votes := s.DBPool.Exec(ctx, sql_votes, input.Image.FileID)
	if err_votes != nil {
		return nil, err_votes
	}

	return &api.Empty{}, nil
}

func (s *Server) GetAllImages(ctx context.Context, page *api.Page) (*api.Images, error) {
	// HW-2 requirement
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
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

	return result, nil
}

func (s *Server) GetGalleryImage(ctx context.Context, request *api.GalleryRequest) (*api.GalleryImage, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	result := &api.GalleryImage{Image: &api.Image{}}

	// get number of images in the gallry
	sql_count := `SELECT
	COUNT(*) AS count
	FROM images
	WHERE userid = $1`
	row_count := s.DBPool.QueryRow(ctx, sql_count, request.UserID)
	err_count := row_count.Scan(&result.Total)
	if err_count != nil {
		return nil, err_count
	}

	if result.Total == 0 {
		return nil, apierr.ErrNoImages
	}

	result.Offset = request.Offset
	if result.Offset >= result.Total {
		result.Offset = result.Total - 1
	}

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

	row_get := s.DBPool.QueryRow(ctx, sql_get, request.UserID, result.Offset)
	err_get := row_get.Scan(&result.Image.FileID, &result.Image.Description, &result.Image.Upvotes, &result.Image.Downvotes)

	if err_get != nil {
		return nil, err_get
	}

	return result, nil
}
