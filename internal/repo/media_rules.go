package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/repo/sqldb"
	"github.com/hs-zavet/tokens/roles"
)

type MediaRulesModel struct {
	ID           string
	Extensions   []string
	MaxSize      int64
	AllowedRoles []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type MediaRulesSQL interface {
	New() sqldb.MediaRulesQ

	Insert(ctx context.Context, input sqldb.MediaRulesInsertInput) (sqldb.MediaRulesModel, error)
	Update(ctx context.Context, input sqldb.MediaRulesUpdateInput) error
	Get(ctx context.Context) (sqldb.MediaRulesModel, error)
	Select(ctx context.Context) ([]sqldb.MediaRulesModel, error)
	Delete(ctx context.Context) error

	FilterID(id string) sqldb.MediaRulesQ

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
	ID           string
	Extensions   []string
	MaxSize      int64
	AllowedRoles []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

func (r MediaRulesRepo) Create(ctx context.Context, input CreateMediaRulesInput) (MediaRulesModel, error) {
	values := sqldb.MediaRulesInsertInput{
		ID:           input.ID,
		Extensions:   input.Extensions,
		MaxSize:      input.MaxSize,
		AllowedRoles: input.AllowedRoles,
		UpdatedAt:    input.UpdatedAt,
		CreatedAt:    input.CreatedAt,
	}

	res, err := r.sql.Insert(ctx, values)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return createMediaRulesModel(res)
}

type MediaRulesUpdateInput struct {
	Extensions   *[]string
	MaxSize      *int64
	AllowedRoles *[]roles.Role
	UpdatedAt    time.Time
}

func (r MediaRulesRepo) Update(ctx context.Context, ID string, input MediaRulesUpdateInput) error {
	var values sqldb.MediaRulesUpdateInput
	if input.Extensions != nil {
		values.Extensions = input.Extensions
	}
	if input.MaxSize != nil {
		values.MaxSize = input.MaxSize
	}
	if input.AllowedRoles != nil {
		values.AllowedRoles = input.AllowedRoles
	}
	values.UpdatedAt = input.UpdatedAt

	return r.sql.New().FilterID(ID).Update(ctx, values)
}

func (r MediaRulesRepo) Get(ctx context.Context, ID string) (MediaRulesModel, error) {
	res, err := r.sql.New().FilterID(ID).Get(ctx)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return createMediaRulesModel(res)
}

func (r MediaRulesRepo) Delete(ctx context.Context, ID string) error {
	err := r.sql.New().FilterID(ID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func createMediaRulesModel(sqlRes sqldb.MediaRulesModel) (MediaRulesModel, error) {
	return MediaRulesModel{
		ID:           sqlRes.ID,
		Extensions:   sqlRes.Extensions,
		MaxSize:      sqlRes.MaxSize,
		AllowedRoles: sqlRes.AllowedRoles,
		UpdatedAt:    sqlRes.UpdatedAt,
		CreatedAt:    sqlRes.CreatedAt,
	}, nil
}
