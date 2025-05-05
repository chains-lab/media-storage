package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const mediaTable = "media"

type MediaModel struct {
	ID         uuid.UUID `db:"id"`
	Format     string    `db:"format"`
	Extensions string    `db:"extensions"`
	Size       int64     `db:"size"`
	Url        string    `db:"url"`
	Resource   string    `db:"resource"`
	ResourceID string    `db:"resource_id"`
	Category   string    `db:"category"`
	OwnerID    uuid.UUID `db:"owner_id,omitempty"`
	CreatedAt  time.Time `db:"created_at"`
}

type MediaQ struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewMedia(db *sql.DB) MediaQ {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return MediaQ{
		db:       db,
		selector: builder.Select("*").From(mediaTable),
		inserter: builder.Insert(mediaTable),
		updater:  builder.Update(mediaTable),
		deleter:  builder.Delete(mediaTable),
		counter:  builder.Select("COUNT(*) AS count").From(mediaTable),
	}
}

func (q MediaQ) New() MediaQ {
	return NewMedia(q.db)
}

type MediaInsertInput struct {
	ID         uuid.UUID `db:"id"`
	Format     string    `db:"format"`
	Extension  string    `db:"extension"`
	Size       int64     `db:"size"`
	Url        string    `db:"url"`
	Resource   string    `db:"resource"`
	ResourceID string    `db:"resource_id"`
	Category   string    `db:"category"`
	OwnerID    uuid.UUID `db:"owner_id,omitempty"`
	CreatedAt  time.Time `db:"created_at"`
}

func (q MediaQ) Insert(ctx context.Context, input MediaInsertInput) (MediaModel, error) {
	values := map[string]any{
		"id":          input.ID,
		"format":      input.Format,
		"extension":   input.Extension,
		"size":        input.Size,
		"url":         input.Url,
		"resource":    input.Resource,
		"resource_id": input.ResourceID,
		"category":    input.Category,
		"owner_id":    input.OwnerID,
		"created_at":  input.CreatedAt,
	}

	query, args, err := q.inserter.SetMap(values).ToSql()
	if err != nil {
		return MediaModel{}, err
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	if err != nil {
		return MediaModel{}, err
	}

	res := MediaModel{
		ID:         input.ID,
		Format:     input.Format,
		Extensions: input.Extension,
		Size:       input.Size,
		Url:        input.Url,
		Resource:   input.Resource,
		ResourceID: input.ResourceID,
		Category:   input.Category,
		OwnerID:    input.OwnerID,
		CreatedAt:  input.CreatedAt,
	}

	return res, err
}

func (q MediaQ) Delete(ctx context.Context) error {
	query, args, err := q.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for accounts: %w", err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q MediaQ) Get(ctx context.Context) (MediaModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return MediaModel{}, err
	}

	var m MediaModel
	err = q.db.QueryRowContext(ctx, query, args...).Scan(
		&m.ID,
		&m.Format,
		&m.Extensions,
		&m.Size,
		&m.Url,
		&m.Resource,
		&m.ResourceID,
		&m.Category,
		&m.OwnerID,
		&m.CreatedAt,
	)
	if err != nil {
		return MediaModel{}, err
	}

	return m, nil
}

func (q MediaQ) Select(ctx context.Context) ([]MediaModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var media []MediaModel
	for rows.Next() {
		var m MediaModel
		err = rows.Scan(
			&m.ID,
			&m.Format,
			&m.Extensions,
			&m.Size,
			&m.Url,
			&m.Resource,
			&m.ResourceID,
			&m.Category,
			&m.OwnerID,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		media = append(media, m)
	}

	return media, nil
}

func (q MediaQ) FilterID(id uuid.UUID) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"id": id})
	q.counter = q.counter.Where(sq.Eq{"id": id})
	q.deleter = q.deleter.Where(sq.Eq{"id": id})
	q.updater = q.updater.Where(sq.Eq{"id": id})
	return q
}

func (q MediaQ) FilterResource(resource string) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"resource": resource})
	q.counter = q.counter.Where(sq.Eq{"resource": resource})
	q.deleter = q.deleter.Where(sq.Eq{"resource": resource})
	q.updater = q.updater.Where(sq.Eq{"resource": resource})
	return q
}

func (q MediaQ) FilterResourceID(resourceID string) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"resource_id": resourceID})
	q.counter = q.counter.Where(sq.Eq{"resource_id": resourceID})
	q.deleter = q.deleter.Where(sq.Eq{"resource_id": resourceID})
	q.updater = q.updater.Where(sq.Eq{"resource_id": resourceID})
	return q
}

func (q MediaQ) FilterCategory(category string) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"category": category})
	q.counter = q.counter.Where(sq.Eq{"category": category})
	q.deleter = q.deleter.Where(sq.Eq{"category": category})
	q.updater = q.updater.Where(sq.Eq{"category": category})
	return q
}

func (q MediaQ) FilterOwnerID(ownerID uuid.UUID) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"owner_id": ownerID})
	q.counter = q.counter.Where(sq.Eq{"owner_id": ownerID})
	q.deleter = q.deleter.Where(sq.Eq{"owner_id": ownerID})
	q.updater = q.updater.Where(sq.Eq{"owner_id": ownerID})
	return q
}

func (q MediaQ) FilterByID(id uuid.UUID) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"id": id})
	q.counter = q.counter.Where(sq.Eq{"id": id})
	q.deleter = q.deleter.Where(sq.Eq{"id": id})
	q.updater = q.updater.Where(sq.Eq{"id": id})
	return q
}

func (q MediaQ) FilterByUrl(url string) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"url": url})
	q.counter = q.counter.Where(sq.Eq{"url": url})
	q.deleter = q.deleter.Where(sq.Eq{"url": url})
	q.updater = q.updater.Where(sq.Eq{"url": url})
	return q
}

func (q MediaQ) Transaction(fn func(ctx context.Context) error) error {
	ctx := context.Background()

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	ctxWithTx := context.WithValue(ctx, txKey, tx)

	if err := fn(ctxWithTx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction failed: %v, rollback error: %v", err, rbErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (q MediaQ) Count(ctx context.Context) (int, error) {
	query, args, err := q.counter.ToSql()
	if err != nil {
		return 0, err
	}

	var count int
	err = q.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q MediaQ) Page(limit, offset uint) MediaQ {
	q.counter = q.counter.Limit(uint64(limit)).Offset(uint64(offset))
	q.selector = q.selector.Limit(uint64(limit)).Offset(uint64(offset))
	return q
}
