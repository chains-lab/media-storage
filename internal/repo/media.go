package repo

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/media-storage/internal/repo/aws"
	"github.com/hs-zavet/media-storage/internal/repo/sqldb"
)

const (
	dataCtxTimeAisle = 10 * time.Second
)

type MediaModel struct {
	ID           uuid.UUID          `db:"id"`
	Folder       string             `db:"folder"`
	Ext          string             `db:"extension"`
	Size         int64              `db:"size"`
	URL          string             `db:"url"`
	ResourceType enums.ResourceType `db:"resource_type"`
	ResourceID   uuid.UUID          `db:"resource_id"`
	MediaType    enums.MediaType    `db:"media_type"`
	OwnerID      uuid.UUID          `db:"owner_id"`
	CreatedAt    time.Time          `db:"created_at"`
}

type mediaAws interface {
	AddFile(ctx context.Context, data aws.FileData, reader io.Reader) (aws.MediaModel, error)
	GetFile(ctx context.Context, data aws.FileData) (aws.MediaModel, error)
	DeleteFile(ctx context.Context, data aws.FileData) error
	DeleteFilesInFolder(ctx context.Context, folder string) error

	ListFiles(ctx context.Context, folder string, offset, limit uint) ([]aws.MediaModel, error)
}

type mediaSQL interface {
	New() sqldb.MediaQ
	Insert(ctx context.Context, input sqldb.MediaInsertInput) (sqldb.MediaModel, error)
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]sqldb.MediaModel, error)
	Get(ctx context.Context) (sqldb.MediaModel, error)

	FilterFilename(id uuid.UUID) sqldb.MediaQ
	FilterFolder(folder string) sqldb.MediaQ

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
	Folder       string
	Filename     uuid.UUID
	Ext          string
	ResourceType enums.ResourceType
	ResourceID   uuid.UUID
	MediaType    enums.MediaType
	OwnerID      uuid.UUID
	CreatedAt    time.Time
}

func (r MediaRepo) AddMedia(ctx context.Context, reader io.Reader, input AddMediaInput) (MediaModel, error) {
	sqlInput := sqldb.MediaInsertInput{
		Filename:     input.Filename,
		Folder:       input.Folder,
		Ext:          input.Ext,
		ResourceType: input.ResourceType,
		ResourceID:   input.ResourceID,
		MediaType:    input.MediaType,
		CreatedAt:    input.CreatedAt,
		OwnerID:      input.OwnerID,
	}

	resSql, err := r.sql.New().Insert(ctx, sqlInput)
	if err != nil {
		return MediaModel{}, fmt.Errorf("sql insert failed: %w", err)
	}

	resAsw, err := r.s3.AddFile(ctx, aws.FileData{
		Folder:   input.Folder,
		Filename: input.Filename,
		Ext:      input.Ext,
	}, reader)
	if err != nil {
		return MediaModel{}, fmt.Errorf("s3 upload failed: %w", err)
	}

	return createMediaModel(resSql, resAsw), nil
}

func (r MediaRepo) GetMedia(ctx context.Context, filename uuid.UUID) (MediaModel, error) {
	sqlMedia, err := r.sql.New().FilterFilename(filename).Get(ctx)
	if err != nil {
		return MediaModel{}, fmt.Errorf("sql get: %w", err)
	}

	s3Media, err := r.s3.ListFiles(ctx, sqlMedia.Folder, 0, 1)
	if err != nil {
		return MediaModel{}, fmt.Errorf("s3 list: %w", err)
	}

	if len(s3Media) == 0 {
		return MediaModel{}, fmt.Errorf("media not found")
	}

	return createMediaModel(sqlMedia, s3Media[0]), nil
}

func (r MediaRepo) DeleteMedia(ctx context.Context, fileId uuid.UUID) error {
	media, err := r.GetMedia(ctx, fileId)
	if err != nil {
		return fmt.Errorf("get media: %w", err)
	}

	err = r.s3.DeleteFile(ctx, aws.FileData{
		Folder:   media.Folder,
		Filename: fileId,
		Ext:      media.Ext,
	})
	if err != nil {
		return fmt.Errorf("s3 delete: %w", err)
	}

	err = r.sql.New().FilterFilename(fileId).Delete(ctx)
	if err != nil {
		return fmt.Errorf("sql delete: %w", err)
	}

	return nil
}

//TODO

func (r MediaRepo) ListMedia(ctx context.Context, folder string, limit, offset uint) ([]MediaModel, error) {
	sqlList, err := r.sql.New().Page(limit, offset).FilterFolder(folder).Select(ctx)
	if err != nil {
		return nil, fmt.Errorf("sql select: %w", err)
	}

	s3List, err := r.s3.ListFiles(ctx, folder, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("s3 list: %w", err)
	}

	results := make([]MediaModel, 0, len(sqlList))
	for i := range sqlList {
		results = append(results, createMediaModel(sqlList[i], s3List[i]))

	}
	return results, nil
}

func (r MediaRepo) DeleteFromFolder(ctx context.Context, folder string) error {
	err := r.s3.DeleteFilesInFolder(ctx, folder)
	if err != nil {
		return fmt.Errorf("s3 batch delete: %w", err)
	}

	err = r.sql.New().FilterFolder(folder).Delete(ctx)
	if err != nil {
		return fmt.Errorf("sql delete folder: %w", err)
	}
	return nil
}

func createMediaModel(sql sqldb.MediaModel, aws aws.MediaModel) MediaModel {
	res := MediaModel{
		ID:           sql.Filename,
		Folder:       sql.Folder,
		Ext:          sql.Ext,
		Size:         aws.Size,
		URL:          aws.URL,
		OwnerID:      sql.OwnerID,
		ResourceType: sql.ResourceType,
		ResourceID:   sql.ResourceID,
		MediaType:    sql.MediaType,
		CreatedAt:    sql.CreatedAt,
	}

	return res
}
