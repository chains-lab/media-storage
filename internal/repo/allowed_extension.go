package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/repo/sqldb"
)

type AllowedExtensionModel struct {
	Resource  string    `db:"resource"`
	Category  string    `db:"category"`
	Extension string    `db:"extension"`
	MaxSize   int64     `db:"max_size"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type allowedExtensionSql interface {
	New() sqldb.AllowedExtensionQ
	Get(ctx context.Context) (sqldb.AllowedExtensionModel, error)
	Select(ctx context.Context) ([]sqldb.AllowedExtensionModel, error)
	FilterByResource(resource string) sqldb.AllowedExtensionQ
	FilterByCategory(category string) sqldb.AllowedExtensionQ
	FilterByExtension(extension string) sqldb.AllowedExtensionQ
}

type AllowedExtension struct {
	sql sqldb.AllowedExtensionQ
}

func NewAllowedExtension(cfg config.Config) (AllowedExtension, error) {
	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return AllowedExtension{}, err
	}
	sqlImpl := sqldb.NewAllowedExtension(db)

	return AllowedExtension{
		sql: sqlImpl,
	}, nil
}

func (a AllowedExtension) GetByResourcesAndCategory(ctx context.Context, resource, category string) ([]sqldb.AllowedExtensionModel, error) {
	allowedExt := a.sql.New().
		FilterByResource(resource).
		FilterByCategory(category)

	return allowedExt.Select(ctx)
}
