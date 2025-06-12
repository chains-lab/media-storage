package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/chains-lab/gatekit/roles"
)

const TableMediaRules = "media_rules"

type MediaRulesModel struct {
	Resource  string    `db:"resource"`
	Category  string    `db:"category"`
	CreatedAt time.Time `db:"created_at"`
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
	ID           string
	Extensions   []string
	MaxSize      int64
	AllowedRoles []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
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

	var rules []MediaRulesModel
	for rows.Next() {
		var r MediaRulesModel
		err = rows.Scan(
			&r.Resource,
			&r.Category,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, r)
	}
	return rules, nil
}

func (q MediaRulesQ) Get(ctx context.Context) (MediaRulesModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return MediaRulesModel{}, err
	}

	var r MediaRulesModel
	err = q.db.QueryRowContext(ctx, query, args...).Scan(
		&r.Resource,
		&r.Category,
		&r.CreatedAt,
	)
	if err != nil {
		return MediaRulesModel{}, err
	}

	return r, nil
}

func (q MediaRulesQ) FilterByResource(resources string) MediaRulesQ {
	q.selector = q.selector.Where(sq.Eq{"resources": resources})
	q.counter = q.counter.Where(sq.Eq{"resources": resources})
	q.deleter = q.deleter.Where(sq.Eq{"resources": resources})
	q.updater = q.updater.Where(sq.Eq{"resources": resources})
	return q
}

func (q MediaRulesQ) FilterByCategory(category string) MediaRulesQ {
	q.selector = q.selector.Where(sq.Eq{"category": category})
	q.counter = q.counter.Where(sq.Eq{"category": category})
	q.deleter = q.deleter.Where(sq.Eq{"category": category})
	q.updater = q.updater.Where(sq.Eq{"category": category})
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
