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
	ID              uuid.UUID          `db:"id"`
	Folder          string             `db:"folder"`
	Ext             string             `db:"extension"`
	ResourceType    enums.ResourceType `db:"resource_type"`
	ResourceID      uuid.UUID          `db:"resource_id"`
	ContentType     enums.ContentType  `db:"content_type"`
	OwnerID         *uuid.UUID         `db:"owner_id,omitempty"`
	Public          bool               `db:"public"`
	AdminOnlyUpdate bool               `db:"admin_only_update"`
	CreatedAt       time.Time          `db:"created_at"`
}

type mediaAws interface {
	AddFile(ctx context.Context, folder, filename, ext string, input aws.AddFileInput) (aws.ContentModel, error)
	ListFiles(ctx context.Context, folder string, offset, limit uint) ([]aws.ContentModel, error)
	DeleteFile(ctx context.Context, folder, filename, ext string) error
	DeleteFilesInFolder(ctx context.Context, folder string) error
}

type mediaSQL interface {
	New() sqldb.MediaQ
	Insert(ctx context.Context, input sqldb.MediaInsertInput) (sqldb.MediaModel, error)
	Delete(ctx context.Context) error

	Select(ctx context.Context) ([]sqldb.MediaModel, error)
	Get(ctx context.Context) (sqldb.MediaModel, error)

	FilterID(id uuid.UUID) sqldb.MediaQ
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
	s3, err := aws.NewS3Client(cfg.Aws.BucketName, cfg.Aws.Region)
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

type AddContentInput struct {
	Reader          io.Reader
	ResourceType    enums.ResourceType
	ResourceID      uuid.UUID
	ContentType     enums.ContentType
	OwnerID         *uuid.UUID
	Public          bool
	AdminOnlyUpdate bool
	CreatedAt       time.Time
}

func (r MediaRepo) AddMedia(ctx context.Context, folder string, filename uuid.UUID, ext string, input AddContentInput) (MediaModel, error) {
	awsInput := aws.AddFileInput{
		Reader: input.Reader,
	}

	resAsw, err := r.s3.AddFile(ctx, folder, filename.String(), ext, awsInput)
	if err != nil {
		return MediaModel{}, fmt.Errorf("s3 upload failed: %w", err)
	}

	inserted := sqldb.MediaInsertInput{
		ID:              filename,
		Folder:          folder,
		Ext:             ext,
		ResourceType:    input.ResourceType,
		ResourceID:      input.ResourceID,
		ContentType:     input.ContentType,
		Public:          input.Public,
		AdminOnlyUpdate: input.AdminOnlyUpdate,
		CreatedAt:       input.CreatedAt,
	}
	if input.OwnerID != nil {
		inserted.OwnerID = input.OwnerID
	}

	resSql, err := r.sql.Insert(ctx, inserted)
	if err != nil {
		return MediaModel{}, fmt.Errorf("sql insert failed: %w", err)
	}

	return createMediaModel(resSql, resAsw), nil
}

func (r MediaRepo) ListMedia(ctx context.Context, folder string, limit, offset uint) ([]MediaModel, error) {
	sqlList, err := r.sql.Page(limit, offset).FilterFolder(folder).Select(ctx)
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

func (r MediaRepo) DeleteMedia(ctx context.Context, folder string, id uuid.UUID, ext string) error {
	err := r.s3.DeleteFile(ctx, folder, id.String(), ext)
	if err != nil {
		return fmt.Errorf("s3 delete: %w", err)
	}

	err = r.sql.New().FilterFolder(folder).FilterID(id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("sql delete: %w", err)
	}

	return nil
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

func createMediaModel(sql sqldb.MediaModel, aws aws.ContentModel) MediaModel {
	res := MediaModel{
		ID:              sql.ID,
		Folder:          sql.Folder,
		Ext:             sql.Ext,
		ResourceType:    sql.ResourceType,
		ResourceID:      sql.ResourceID,
		ContentType:     sql.ContentType,
		Public:          sql.Public,
		AdminOnlyUpdate: sql.AdminOnlyUpdate,
		CreatedAt:       sql.CreatedAt,
	}

	if sql.OwnerID != nil {
		res.OwnerID = sql.OwnerID
	}

	return res
}
