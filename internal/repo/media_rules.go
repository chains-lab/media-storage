package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/repo/sqldb"
)

const (
	dataCtxTimeAisle = 10 * time.Second
)

type MediaModel struct {
	Resource  string
	Category  string
	CreatedAt time.Time
}

type mediaRulesSql interface {
	New() sqldb.MediaRulesQ
	Select(ctx context.Context) ([]sqldb.MediaRulesModel, error)
	Get(ctx context.Context) (sqldb.MediaRulesModel, error)
	Count(ctx context.Context) (int, error)
	FilterByResource(resource string) sqldb.MediaRulesQ
	FilterByCategory(category string) sqldb.MediaRulesQ
	Page(limit, offset uint) sqldb.MediaRulesQ
}

type Media struct {
	sql sqldb.MediaRulesQ
}

func NewMediaRules(cfg config.Config) (Media, error) {
	db, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		return Media{}, err
	}
	sqlImpl := sqldb.NewMediaRules(db)

	return Media{
		sql: sqlImpl,
	}, nil
}

func (m Media) GetByResourceAndCategory(ctx context.Context, resource, category string) (sqldb.MediaRulesModel, error) {
	mediaRules := m.sql.New().
		FilterByResource(resource).
		FilterByCategory(category)

	return mediaRules.Get(ctx)
}
