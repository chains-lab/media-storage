package repo

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/repo/aws"
	"github.com/chains-lab/media-storage/internal/repo/sqldb"
	"github.com/google/uuid"
)

const (
	dataCtxTimeAisle = 10 * time.Second
)

type MediaModel struct {
	Resource  string    `db:"resource"`
	Category  string    `db:"category"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

type mediaAws interface {
	GeneratePutURL(ctx context.Context, folder, originalFilename, contentType string) (*aws.MediaURL, error)
	GenerateGetURL(ctx context.Context, key string) (string, error)
}

type mediaSQL interface {
	New() sqldb.MediaQ
	Insert(ctx context.Context, input sqldb.MediaInsertInput) (sqldb.MediaModel, error)
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]sqldb.MediaModel, error)
	Get(ctx context.Context) (sqldb.MediaModel, error)

	FilterID(id uuid.UUID) sqldb.MediaQ
	FilterResource(resourceType string) sqldb.MediaQ
	FilterResourceID(resourceID string) sqldb.MediaQ
	FilterCategory(category string) sqldb.MediaQ
	FilterOwnerID(ownerID uuid.UUID) sqldb.MediaQ
	FilterByID(id uuid.UUID) sqldb.MediaQ
	FilterByUrl(url string) sqldb.MediaQ

	Count(ctx context.Context) (int, error)
	Transaction(fn func(ctx context.Context) error) error
	Page(limit, offset uint) sqldb.MediaQ
}

type MediaRepo struct {
	s3  mediaAws
	sql mediaSQL
}

func NewMedia(cfg config.Config) (MediaRepo, error) {
	s3, err := aws.NewAwsS3Client(cfg.Aws.BucketName, cfg.Aws.Region, cfg.Aws.AccessKeyID, cfg.Aws.AccessKey)
	if err != nil {
		return MediaRepo{}, err
	}

	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return MediaRepo{}, err
	}

	return MediaRepo{
		s3:  s3,
		sql: sqldb.NewMedia(db),
	}, nil
}

type AddMediaInput struct {
	Filename   string // Filename is original filename without changes and without path
	Resource   string
	ResourceID string
	Category   string
	OwnerID    uuid.UUID
	CreatedAt  time.Time
}

func (r MediaRepo) UploadMedia(ctx context.Context, reader io.Reader, input AddMediaInput) (MediaModel, error) {
	resAsw, err := r.s3.UploadFile(ctx, reader, fmt.Sprintf("%s-%s", input.Resource, input.Category), input.Filename)
	if err != nil {
		return MediaModel{}, fmt.Errorf("s3 upload failed: %w", err)
	}

	id, err := uuid.Parse(resAsw.ID)
	if err != nil {
		return MediaModel{}, fmt.Errorf("uuid parse failed: %w", err)
	}

	sqlInput := sqldb.MediaInsertInput{
		ID:         id,
		Format:     resAsw.ContentType,
		Extension:  resAsw.Extension,
		Size:       resAsw.Size,
		Url:        resAsw.Url,
		Resource:   input.Resource,
		ResourceID: input.ResourceID,
		Category:   input.Category,
		OwnerID:    input.OwnerID,
		CreatedAt:  input.CreatedAt,
	}

	resSql, err := r.sql.Insert(ctx, sqlInput)
	if err != nil {
		return MediaModel{}, fmt.Errorf("sql insert failed: %w", err)
	}

	return createMediaModel(resSql, resAsw), nil
}

func (r MediaRepo) GetMedia(ctx context.Context, mediaID uuid.UUID) (MediaModel, error) {
	sqlMedia, err := r.sql.New().FilterID(mediaID).Get(ctx)
	if err != nil {
		return MediaModel{}, fmt.Errorf("sql get: %w", err)
	}

	s3Media, err := r.s3.GetFileByUrl(ctx, sqlMedia.Url)
	if err != nil {
		return MediaModel{}, fmt.Errorf("s3 list: %w", err)
	}

	return createMediaModel(sqlMedia, s3Media), nil
}

func (r MediaRepo) DeleteMedia(ctx context.Context, mediaID uuid.UUID) error {
	media, err := r.GetMedia(ctx, mediaID)
	if err != nil {
		return fmt.Errorf("get media: %w", err)
	}

	err = r.s3.DeleteFileByUrl(ctx, media.Url)
	if err != nil {
		return fmt.Errorf("s3 delete: %w", err)
	}

	err = r.sql.New().FilterID(mediaID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("sql delete: %w", err)
	}

	return nil
}

func (r MediaRepo) DeleteFilesByResourceAndCategory(ctx context.Context, resource, category string) error {
	err := r.s3.DeleteFilesByPrefix(ctx, fmt.Sprintf("%s-%s", resource, category))
	if err != nil {
		return fmt.Errorf("s3 delete by prefix: %w", err)
	}

	err = r.sql.New().FilterResource(resource).FilterCategory(category).Delete(ctx)
	if err != nil {
		return fmt.Errorf("sql delete by prefix: %w", err)
	}

	return nil
}

func createMediaModel(sql sqldb.MediaModel, aws aws.MediaModel) MediaModel {
	res := MediaModel{
		ID:         sql.ID,
		Format:     aws.ContentType,
		Extension:  aws.Extension,
		Size:       aws.Size,
		Url:        aws.Url,
		Resource:   sql.Resource,
		ResourceID: sql.ResourceID,
		Category:   sql.Category,
		OwnerID:    sql.OwnerID,
		CreatedAt:  sql.CreatedAt,
	}

	return res
}
