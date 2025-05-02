package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

const TableMediaRules = "media_rules"

type MediaRulesModel struct {
	MediaType    enums.MediaType `db:"media_type"`
	MaxSize      int64           `db:"max_size"`
	AllowedExits []string        `db:"allowed_exits"`
	Folder       string          `db:"folder"`
	Roles        []roles.Role    `db:"roles_access_update"`
}

type MediaRulesQ struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewMediaRules(db *sql.DB) MediaRulesQ {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return MediaRulesQ{
		db:       db,
		selector: builder.Select("*").From(TableMediaRules),
		inserter: builder.Insert(TableMediaRules),
		updater:  builder.Update(TableMediaRules),
		deleter:  builder.Delete(TableMediaRules),
		counter:  builder.Select("COUNT(*) AS count").From(TableMediaRules),
	}
}

func (q MediaRulesQ) New() MediaRulesQ {
	return NewMediaRules(q.db)
}

type MediaRulesInsertInput struct {
	MediaType    enums.MediaType
	MaxSize      int64
	AllowedExits []string
	Folder       string
	Roles        []roles.Role
}

func (q MediaRulesQ) Insert(ctx context.Context, input MediaRulesInsertInput) (MediaRulesModel, error) {
	values := map[string]interface{}{
		"media_type":          input.MediaType,
		"max_size":            input.MaxSize,
		"allowed_exits":       input.AllowedExits,
		"folder":              input.Folder,
		"roles_access_update": input.Roles,
	}

	query, args, err := q.inserter.Values(values).ToSql()
	if err != nil {
		return MediaRulesModel{}, err
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	if err != nil {
		return MediaRulesModel{}, err
	}

	res := MediaRulesModel{
		MediaType:    input.MediaType,
		MaxSize:      input.MaxSize,
		AllowedExits: input.AllowedExits,
		Folder:       input.Folder,
		Roles:        input.Roles,
	}

	return res, nil
}

type MediaRulesUpdateInput struct {
	MaxSize      *int64
	AllowedExits *[]string
	Folder       *string
	Roles        *[]roles.Role
}

func (q MediaRulesQ) Update(ctx context.Context, input MediaRulesUpdateInput) error {
	values := map[string]interface{}{}
	if input.MaxSize != nil {
		values["max_size"] = *input.MaxSize
	}
	if input.AllowedExits != nil {
		values["allowed_exits"] = *input.AllowedExits
	}
	if input.Folder != nil {
		values["folder"] = *input.Folder
	}
	if input.Roles != nil {
		values["roles_access_update"] = *input.Roles
	}

	query, args, err := q.updater.SetMap(values).ToSql()
	if err != nil {
		return err
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q MediaRulesQ) Delete(ctx context.Context) error {
	query, args, err := q.deleter.ToSql()
	if err != nil {
		return err
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q MediaRulesQ) Select(ctx context.Context) ([]MediaRulesModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaRules []MediaRulesModel
	for rows.Next() {
		var rule MediaRulesModel
		err = rows.Scan(
			&rule.MediaType,
			&rule.MaxSize,
			&rule.AllowedExits,
			&rule.Folder,
			&rule.Roles,
		)
		if err != nil {
			return nil, err
		}
		mediaRules = append(mediaRules, rule)
	}
	return mediaRules, nil
}

func (q MediaRulesQ) Get(ctx context.Context) (MediaRulesModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return MediaRulesModel{}, err
	}

	var rule MediaRulesModel
	err = q.db.QueryRowContext(ctx, query, args...).Scan(
		&rule.MediaType,
		&rule.MaxSize,
		&rule.AllowedExits,
		&rule.Folder,
		&rule.Roles,
	)
	if err != nil {
		return MediaRulesModel{}, err
	}
	return rule, nil
}

func (q MediaRulesQ) FilterMediaType(mediaType enums.MediaType) MediaRulesQ {
	q.selector = q.selector.Where(sq.Eq{"media_type": mediaType})
	q.counter = q.counter.Where(sq.Eq{"media_type": mediaType})
	q.deleter = q.deleter.Where(sq.Eq{"media_type": mediaType})
	q.updater = q.updater.Where(sq.Eq{"media_type": mediaType})
	return q
}

func (q MediaRulesQ) FilterFolder(folder string) MediaRulesQ {
	q.selector = q.selector.Where(sq.Eq{"folder": folder})
	q.counter = q.counter.Where(sq.Eq{"folder": folder})
	q.deleter = q.deleter.Where(sq.Eq{"folder": folder})
	q.updater = q.updater.Where(sq.Eq{"folder": folder})
	return q
}

func (q MediaRulesQ) Transaction(fn func(ctx context.Context) error) error {
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

func (q MediaRulesQ) Count(ctx context.Context) (int, error) {
	query, args, err := q.counter.ToSql()
	if err != nil {
		return 0, err
	}

	var count int
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		err = tx.QueryRowContext(ctx, query, args...).Scan(&count)
	} else {
		err = q.db.QueryRowContext(ctx, query, args...).Scan(&count)
	}

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (q MediaRulesQ) Page(limit, offset uint) MediaRulesQ {
	q.counter = q.counter.Limit(uint64(limit)).Offset(uint64(offset))
	q.selector = q.selector.Limit(uint64(limit)).Offset(uint64(offset))
	return q
}
