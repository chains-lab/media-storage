package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/enums"
)

const mediaTable = "media"

type MediaModel struct {
	ID     uuid.UUID `db:"id"`
	Folder string    `db:"folder"`
	Ext    string    `db:"ext"`

	//Name of resource
	ResourceType enums.ResourceType `db:"resource_type"`
	//ID of resource
	ResourceID uuid.UUID `db:"resource_id"`

	//MediaType of resource
	MediaType enums.MediaType `db:"media_type"`

	//Owner ID of resource who
	OwnerID *uuid.UUID `db:"owner_id,omitempty"`
	//Public          bool       `db:"public"`
	//AdminOnlyUpdate bool       `db:"admin_only_update"`

	CreatedAt time.Time `db:"created_at"`
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
	ID           uuid.UUID          `db:"id"`
	Folder       string             `db:"folder"`
	Ext          string             `db:"extension"`
	ResourceType enums.ResourceType `db:"resource_type"`
	ResourceID   uuid.UUID          `db:"resource_id"`

	MediaType enums.MediaType `db:"media_type"`

	OwnerID *uuid.UUID `db:"owner_id,omitempty"`
	//Public          bool       `db:"public"`
	//AdminOnlyUpdate bool       `db:"admin_only_update"`

	CreatedAt time.Time `db:"created_at"`
}

func (q MediaQ) Insert(ctx context.Context, input MediaInsertInput) (MediaModel, error) {
	values := map[string]any{
		"id":            input.ID,
		"folder":        input.Folder,
		"extension":     input.Ext,
		"resource_type": input.ResourceType,
		"resource_id":   input.ResourceID,
		"media_type":    input.MediaType,
		"created_at":    input.CreatedAt,
	}

	if input.OwnerID != nil {
		values["owner_id"] = *input.OwnerID
	} else {
		values["owner_id"] = nil
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

	res := MediaModel{
		ID:           input.ID,
		Folder:       input.Folder,
		Ext:          input.Ext,
		ResourceType: input.ResourceType,
		ResourceID:   input.ResourceID,
		MediaType:    input.MediaType,
		CreatedAt:    input.CreatedAt,
	}

	if input.OwnerID != nil {
		res.OwnerID = input.OwnerID
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
			&m.ID,
			&m.Folder,
			&m.Ext,
			&m.ResourceType,
			&m.ResourceID,
			&m.MediaType,
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

func (q MediaQ) Get(ctx context.Context) (MediaModel, error) {
	query, args, err := q.deleter.ToSql()
	if err != nil {
		return MediaModel{}, fmt.Errorf("building delete query for accounts: %w", err)
	}

	var m MediaModel
	err = q.db.QueryRowContext(ctx, query, args...).Scan(
		&m.ID,
		&m.Folder,
		&m.Ext,
		&m.ResourceType,
		&m.ResourceID,
		&m.MediaType,
		&m.OwnerID,
		&m.CreatedAt,
	)
	if err != nil {
		return MediaModel{}, err
	}

	return m, nil
}

func (q MediaQ) FilterID(id uuid.UUID) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"id": id})
	q.counter = q.counter.Where(sq.Eq{"id": id})
	q.deleter = q.deleter.Where(sq.Eq{"id": id})
	q.updater = q.updater.Where(sq.Eq{"id": id})
	return q
}

func (q MediaQ) FilterFolder(folder string) MediaQ {
	q.selector = q.selector.Where(sq.Eq{"folder": folder})
	q.counter = q.counter.Where(sq.Eq{"folder": folder})
	q.deleter = q.deleter.Where(sq.Eq{"folder": folder})
	q.updater = q.updater.Where(sq.Eq{"folder": folder})
	return q
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

func (q MediaQ) Page(limit, offset uint) MediaQ {
	q.counter = q.counter.Limit(uint64(limit)).Offset(uint64(offset))
	q.selector = q.selector.Limit(uint64(limit)).Offset(uint64(offset))
	return q
}
