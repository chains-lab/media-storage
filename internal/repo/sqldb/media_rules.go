package sqldb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

const TableMediaRules = "media_rules"

type MediaRulesModel struct {
	ResourceType string           `db:"resource_type"`
	ExitSize     []enums.ExitSize `db:"exit_size"`
	Roles        []roles.Role     `db:"roles_access_update"`
	UpdatedAt    time.Time        `db:"updated_at"`
	CreatedAt    time.Time        `db:"created_at"`
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

type ExitSizeJSON []enums.ExitSize

func (e *ExitSizeJSON) Value() (driver.Value, error) {
	return json.Marshal(e)
}
func (e *ExitSizeJSON) Scan(src any) error {
	if src == nil {
		*e = nil
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("exit_size: expect []byte, got %T", src)
	}
	return json.Unmarshal(b, (*[]enums.ExitSize)(e))
}

func (q MediaRulesQ) New() MediaRulesQ {
	return NewMediaRules(q.db)
}

type RolesJSON []roles.Role

func (r *RolesJSON) Value() (driver.Value, error) {
	return json.Marshal(r)
}
func (r *RolesJSON) Scan(src any) error {
	if src == nil {
		*r = nil
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("roles_access: expect []byte, got %T", src)
	}
	return json.Unmarshal(b, r)
}

type MediaRulesInsertInput struct {
	ResourceType string
	ExitSize     []enums.ExitSize
	Roles        []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

func (q MediaRulesQ) Insert(ctx context.Context, in MediaRulesInsertInput) (MediaRulesModel, error) {
	vals := map[string]any{
		"resource_type": in.ResourceType,
		"exit_size":     ExitSizeJSON(in.ExitSize), // ← JSONB
		"roles_access":  RolesJSON(in.Roles),       // ← JSONB
		"updated_at":    in.UpdatedAt,
		"created_at":    in.CreatedAt,
	}

	query, args, err := q.inserter.SetMap(vals).ToSql()
	if err != nil {
		return MediaRulesModel{}, err
	}

	executor := q.db.ExecContext
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		executor = tx.ExecContext
	}
	if _, err = executor(ctx, query, args...); err != nil {
		return MediaRulesModel{}, err
	}

	return MediaRulesModel{
		ResourceType: in.ResourceType,
		ExitSize:     in.ExitSize,
		Roles:        in.Roles,
		UpdatedAt:    in.UpdatedAt,
		CreatedAt:    in.CreatedAt,
	}, nil
}

type MediaRulesUpdateInput struct {
	ExitSize  *[]enums.ExitSize
	Roles     *[]roles.Role
	UpdatedAt time.Time
}

func (q MediaRulesQ) Update(ctx context.Context, in MediaRulesUpdateInput) error {
	vals := map[string]any{"updated_at": in.UpdatedAt}

	if in.ExitSize != nil {
		vals["exit_size"] = ExitSizeJSON(*in.ExitSize)
	}
	if in.Roles != nil {
		vals["roles_access"] = RolesJSON(*in.Roles)
	}

	query, args, err := q.updater.SetMap(vals).ToSql()
	if err != nil {
		return err
	}

	executor := q.db.ExecContext
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		executor = tx.ExecContext
	}
	_, err = executor(ctx, query, args...)
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

type scanner interface{ Scan(dest ...any) error }

func scanRule(s scanner, out *MediaRulesModel) error {
	var es ExitSizeJSON
	var rl RolesJSON
	if err := s.Scan(
		&out.ResourceType,
		&es,
		&rl,
		&out.UpdatedAt,
		&out.CreatedAt,
	); err != nil {
		return err
	}
	out.ExitSize = []enums.ExitSize(es)
	out.Roles = []roles.Role(rl)
	return nil
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

	var list []MediaRulesModel
	for rows.Next() {
		var m MediaRulesModel
		if err := scanRule(rows, &m); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}

func (q MediaRulesQ) Get(ctx context.Context) (MediaRulesModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return MediaRulesModel{}, err
	}

	var m MediaRulesModel
	if err := scanRule(q.db.QueryRowContext(ctx, query, args...), &m); err != nil {
		return MediaRulesModel{}, err
	}
	return m, nil
}

func (q MediaRulesQ) FilterResourceType(resourceType string) MediaRulesQ {
	q.selector = q.selector.Where(sq.Eq{"resource_type": resourceType})
	q.counter = q.counter.Where(sq.Eq{"resource_type": resourceType})
	q.deleter = q.deleter.Where(sq.Eq{"resource_type": resourceType})
	q.updater = q.updater.Where(sq.Eq{"resource_type": resourceType})
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
