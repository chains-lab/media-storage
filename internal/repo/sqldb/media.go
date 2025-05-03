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
	Filename     uuid.UUID `db:"filename"`
	Ext          string    `db:"ext"`
	ResourceType string    `db:"resource_type"`
	ResourceID   uuid.UUID `db:"resource_id"`
	OwnerID      uuid.UUID `db:"owner_id,omitempty"`
	UpdatedAt    time.Time `db:"updated_at"`
	CreatedAt    time.Time `db:"created_at"`
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
	Filename     uuid.UUID `db:"filename"`
	Ext          string    `db:"extension"`
	ResourceType string    `db:"resource_type"`
	ResourceID   uuid.UUID `db:"resource_id"`
	OwnerID      uuid.UUID `db:"owner_id,"`
	UpdatedAt    time.Time `db:"updated_at"`
	CreatedAt    time.Time `db:"created_at"`
}

func (q MediaQ) Insert(ctx context.Context, input MediaInsertInput) (MediaModel, error) {
	values := map[string]any{
		"filename":      input.Filename,
		"extension":     input.Ext,
		"resource_type": input.ResourceType,
		"resource_id":   input.ResourceID,
		"owner_id":      input.OwnerID,
		"updated_at":    input.UpdatedAt,
		"created_at":    input.CreatedAt,
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
		Filename:     input.Filename,
		Ext:          input.Ext,
		ResourceType: input.ResourceType,
		ResourceID:   input.ResourceID,
		OwnerID:      input.OwnerID,
		UpdatedAt:    input.UpdatedAt,
		CreatedAt:    input.CreatedAt,
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
	if err != nil {
		return err
	}

	return nil
}

func (q MediaQ) Get(ctx context.Context) (MediaModel, error) {
	query, args, err := q.deleter.ToSql()
	if err != nil {
		return MediaModel{}, fmt.Errorf("building delete query for accounts: %w", err)
	}

	var m MediaModel
	err = q.db.QueryRowContext(ctx, query, args...).Scan(
		&m.Filename,
		&m.Ext,
		&m.ResourceType,
		&m.ResourceID,
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
		err := rows.Scan(
			&m.Filename,
			&m.Ext,
			&m.ResourceType,
			&m.ResourceID,
			&m.OwnerID,
			&m.UpdatedAt,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		media = append(media, m)
	}

	return media, nil
}

func (q MediaQ) FilterFilename(name uuid.UUID) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"filename": name})
	q.counter = q.counter.Where(sq.Eq{"filename": name})
	q.deleter = q.deleter.Where(sq.Eq{"filename": name})
	q.updater = q.updater.Where(sq.Eq{"filename": name})
	return q
}

func (q MediaQ) FilterResourceType(resourceType string) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"resource_type": resourceType})
	q.counter = q.counter.Where(sq.Eq{"resource_type": resourceType})
	q.deleter = q.deleter.Where(sq.Eq{"resource_type": resourceType})
	q.updater = q.updater.Where(sq.Eq{"resource_type": resourceType})
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
