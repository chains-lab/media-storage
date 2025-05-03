package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/media-storage/internal/repo/sqldb"
	"github.com/hs-zavet/tokens/roles"
)

type MediaRulesModel struct {
	ResourceType string           `db:"resource_type"`
	ExitSize     []enums.ExitSize `db:"exit_size"`
	Roles        []roles.Role     `db:"roles_access_update"`
	UpdatedAt    time.Time        `db:"updated_at"`
	CreatedAt    time.Time        `db:"created_at"`
}

type MediaRulesSQL interface {
	New() sqldb.MediaRulesQ

	Insert(ctx context.Context, input sqldb.MediaRulesInsertInput) (sqldb.MediaRulesModel, error)
	Update(ctx context.Context, input sqldb.MediaRulesUpdateInput) error
	Get(ctx context.Context) (sqldb.MediaRulesModel, error)
	Select(ctx context.Context) ([]sqldb.MediaRulesModel, error)
	Delete(ctx context.Context) error

	FilterResourceType(resourceType string) sqldb.MediaRulesQ

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
	ResourceType string           `db:"resource_type"`
	ExitSize     []enums.ExitSize `db:"exit_size"`
	Roles        []roles.Role     `db:"roles_access_update"`
	UpdatedAt    time.Time        `db:"updated_at"`
	CreatedAt    time.Time        `db:"created_at"`
}

func (r MediaRulesRepo) Create(ctx context.Context, input CreateMediaRulesInput) (MediaRulesModel, error) {
	values := sqldb.MediaRulesInsertInput{
		ResourceType: input.ResourceType,
		ExitSize:     input.ExitSize,
		Roles:        input.Roles,
		UpdatedAt:    input.UpdatedAt,
		CreatedAt:    input.CreatedAt,
	}

	res, err := r.sql.Insert(ctx, values)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return createMediaRulesModel(res), nil
}

type MediaRulesUpdateInput struct {
	ExitSize  *[]enums.ExitSize `db:"exit_size"`
	Roles     *[]roles.Role     `db:"roles_access_update"`
	UpdatedAt time.Time         `db:"created_at"`
}

func (r MediaRulesRepo) Update(ctx context.Context, resourceType string, input MediaRulesUpdateInput) error {
	var values sqldb.MediaRulesUpdateInput
	if input.ExitSize != nil {
		values.ExitSize = input.ExitSize
	}
	if input.Roles != nil {
		values.Roles = input.Roles
	}
	values.UpdatedAt = input.UpdatedAt

	return r.sql.Update(ctx, values)
}

func (r MediaRulesRepo) Get(ctx context.Context, resourceType string) (MediaRulesModel, error) {
	res, err := r.sql.New().FilterResourceType(resourceType).Get(ctx)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return createMediaRulesModel(res), nil
}

func (r MediaRulesRepo) Delete(ctx context.Context, resourceType string) error {
	err := r.sql.New().FilterResourceType(resourceType).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func createMediaRulesModel(input sqldb.MediaRulesModel) MediaRulesModel {
	return MediaRulesModel{
		ResourceType: input.ResourceType,
		ExitSize:     input.ExitSize,
		Roles:        input.Roles,
		UpdatedAt:    input.UpdatedAt,
		CreatedAt:    input.CreatedAt,
	}
}
