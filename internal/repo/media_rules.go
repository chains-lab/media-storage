package repo

import (
	"context"
	"database/sql"

	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/media-storage/internal/repo/sqldb"
	"github.com/hs-zavet/tokens/roles"
)

type MediaRulesModel struct {
	MediaType    enums.MediaType `db:"media_type"`
	MaxSize      int64           `db:"max_size"`
	AllowedExits []string        `db:"allowed_exits"`
	Folder       string          `db:"folder"`
	Roles        []roles.Role    `db:"roles_access_update"`
}

type MediaRulesSQL interface {
	New() sqldb.MediaRulesQ

	Insert(ctx context.Context, input sqldb.MediaRulesInsertInput) (sqldb.MediaRulesModel, error)
	Update(ctx context.Context, input sqldb.MediaRulesUpdateInput) error
	Get(ctx context.Context) (sqldb.MediaRulesModel, error)
	Select(ctx context.Context) ([]sqldb.MediaRulesModel, error)
	Delete(ctx context.Context) error

	FilterMediaType(mediaType enums.MediaType) sqldb.MediaRulesQ
	FilterFolder(folder string) sqldb.MediaRulesQ

	Transaction(fn func(ctx context.Context) error) error
	Count(ctx context.Context) (int, error)
	Page(limit, offset uint) sqldb.MediaRulesQ
}

type MediaRulesRepo struct {
	sql MediaRulesSQL
}

func NewMediaRulesRepo(cfg config.Config) (MediaRulesRepo, error) {
	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return MediaRulesRepo{}, err
	}

	return MediaRulesRepo{
		sql: sqldb.NewMediaRules(db),
	}, nil
}

type CreateMediaRulesInput struct {
	MediaType    enums.MediaType
	MaxSize      int64
	AllowedExits []string
	Folder       string
	Roles        []roles.Role
}

func (r MediaRulesRepo) Create(ctx context.Context, input CreateMediaRulesInput) (MediaRulesModel, error) {
	values := sqldb.MediaRulesInsertInput{
		MediaType:    input.MediaType,
		MaxSize:      input.MaxSize,
		AllowedExits: input.AllowedExits,
		Folder:       input.Folder,
		Roles:        input.Roles,
	}

	res, err := r.sql.Insert(ctx, values)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return createMediaRulesModel(res), nil
}

type MediaRulesUpdateInput struct {
	MaxSize      *int64
	AllowedExits *[]string
	Folder       *string
	Roles        *[]roles.Role
}

func (r MediaRulesRepo) Update(ctx context.Context, input MediaRulesUpdateInput) error {
	var values sqldb.MediaRulesUpdateInput
	if input.MaxSize != nil {
		values.MaxSize = input.MaxSize
	}
	if input.AllowedExits != nil {
		values.AllowedExits = input.AllowedExits
	}
	if input.Folder != nil {
		values.Folder = input.Folder
	}
	if input.Roles != nil {
		values.Roles = input.Roles
	}

	return r.sql.Update(ctx, values)
}

func (r MediaRulesRepo) Get(ctx context.Context, mType enums.MediaType) (MediaRulesModel, error) {
	res, err := r.sql.New().FilterMediaType(mType).Get(ctx)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return createMediaRulesModel(res), nil
}

func (r MediaRulesRepo) Delete(ctx context.Context, mType enums.MediaType) error {
	err := r.sql.New().FilterMediaType(mType).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func createMediaRulesModel(input sqldb.MediaRulesModel) MediaRulesModel {
	return MediaRulesModel{
		MediaType:    input.MediaType,
		MaxSize:      input.MaxSize,
		AllowedExits: input.AllowedExits,
		Folder:       input.Folder,
		Roles:        input.Roles,
	}
}
