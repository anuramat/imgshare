package imgshare

import (
	"context"
	"log"

	"gitlab.ozon.dev/anuramat/homework-1/internal/api"
	"gitlab.ozon.dev/anuramat/homework-1/internal/apierr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateImage(ctx context.Context, input *api.ImageAuthRequest) (*api.Empty, error) {
	s.pool <- struct{}{}
	defer func() { <-s.pool }()

	select {
	case <-ctx.Done():
		return nil, apierr.ErrTimeout
	default:
	}

	query := "INSERT INTO images (fileid, userid, description) VALUES ($1, $2, $3);"
	_, err := s.dbpool.Exec(ctx, query, input.Image.FileID, input.UserID, input.Image.Description)
	if err != nil {
		log.Println(err)
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

	query_count := "SELECT COUNT(*) FROM images;"
	row_count := s.dbpool.QueryRow(ctx, query_count)
	err_count := row_count.Scan(&count)
	if err_count != nil {
		log.Println(err_count)
		return nil, err_count
	}

	if count == 0 {
		return nil, status.Error(codes.NotFound, "no images yet")
	}

	query_random := "SELECT fileid FROM images ORDER BY random() LIMIT 1;"
	row_random := s.dbpool.QueryRow(ctx, query_random)
	err_random := row_random.Scan(&fileID)
	if err_random != nil {
		log.Println(err_random)
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

	query := "INSERT INTO votes (fileid, userid, upvote) VALUES ($1, $2, $3) ON CONFLICT (fileid, userid) DO UPDATE SET upvote = $3;"
	_, err := s.dbpool.Exec(ctx, query, fid, uid, vote)
	if err != nil {
		log.Println(err)
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

	query := "UPDATE images SET description = $1 WHERE fileid = $2 AND userid = $3;"
	_, err := s.dbpool.Exec(ctx, query, input.Image.Description, input.Image.FileID, input.UserID)
	if err != nil {
		log.Println(err)
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

	query_images := "DELETE FROM images WHERE fileid = $1 AND userid = $2;"
	res, err_images := s.dbpool.Exec(ctx, query_images, input.Image.FileID, input.UserID)
	if err_images != nil {
		log.Println(err_images)
		return nil, err_images
	}
	count := res.RowsAffected()
	if count == 0 {
		return &api.Empty{}, nil
	}

	query_votes := "DELETE FROM votes WHERE fileid = $1;"
	_, err_votes := s.dbpool.Exec(ctx, query_votes, input.Image.FileID)
	if err_votes != nil {
		log.Println(err_votes)
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

	query := `SELECT
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
	rows, err := s.dbpool.Query(ctx, query, page.Limit, page.Offset)
	if err != nil {
		log.Println(err)
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

	query := `SELECT
	images.fileid as fileid,
	images.description as description,
	COUNT(*) FILTER (WHERE votes.upvote) as upvotes,
	COUNT(*) FILTER (WHERE NOT votes.upvote) as downvotes
	FROM images
	LEFT JOIN votes ON images.fileid = votes.fileid
	WHERE images.fileid = $1
	GROUP BY images.fileid`

	row := s.dbpool.QueryRow(ctx, query, fileID)
	err := row.Scan(&result.FileID, &result.Description, &result.Upvotes, &result.Downvotes)

	if err != nil {
		log.Println(err)
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
	query_count := `SELECT
	COUNT(*) AS count
	FROM images
	WHERE userid = $1`
	row_count := s.dbpool.QueryRow(ctx, query_count, request.UserID)
	err_count := row_count.Scan(&result.Total)
	if err_count != nil {
		log.Println(err_count)
		return nil, err_count
	}

	if result.Total == 0 {
		return nil, status.Error(codes.NotFound, "no images yet")
	}

	result.Offset = request.Offset
	if result.Offset >= result.Total {
		result.Offset = result.Total - 1
	}

	query_get := `SELECT
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

	row_get := s.dbpool.QueryRow(ctx, query_get, request.UserID, result.Offset)
	err_get := row_get.Scan(&result.Image.FileID, &result.Image.Description, &result.Image.Upvotes, &result.Image.Downvotes)

	if err_get != nil {
		log.Println(err_get)
		return nil, err_get
	}

	return result, nil
}
