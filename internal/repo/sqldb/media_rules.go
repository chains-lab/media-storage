package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hs-zavet/tokens/roles"
	"github.com/lib/pq"
)

const TableMediaRules = "media_rules"

type MediaRulesModel struct {
	ID           string       `db:"id"`
	Extensions   []string     `db:"extensions"`
	MaxSize      int64        `db:"max_size"`
	AllowedRoles []roles.Role `db:"allowed_roles"`
	UpdatedAt    time.Time    `db:"updated_at"`
	CreatedAt    time.Time    `db:"created_at"`
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

func (q MediaRulesQ) Insert(ctx context.Context, in MediaRulesInsertInput) (MediaRulesModel, error) {
	rolesInput := make([]string, len(in.AllowedRoles))
	for i, role := range in.AllowedRoles {
		rolesInput[i] = string(role)
	}
	vals := map[string]any{
		"id":            in.ID,
		"extensions":    pq.Array(in.Extensions),
		"allowed_roles": pq.Array(rolesInput),
		"max_size":      in.MaxSize,
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
		ID:           in.ID,
		Extensions:   in.Extensions,
		MaxSize:      in.MaxSize,
		AllowedRoles: in.AllowedRoles,
		UpdatedAt:    in.UpdatedAt,
		CreatedAt:    in.CreatedAt,
	}, err
}

type MediaRulesUpdateInput struct {
	Extensions   *[]string
	MaxSize      *int64
	AllowedRoles *[]roles.Role
	UpdatedAt    time.Time
}

func (q MediaRulesQ) Update(ctx context.Context, in MediaRulesUpdateInput) error {
	vals := map[string]any{"updated_at": in.UpdatedAt}

	if in.Extensions != nil {
		vals["extensions"] = pq.Array(*in.Extensions)
	}
	if in.MaxSize != nil {
		vals["max_size"] = *in.MaxSize
	}
	if in.AllowedRoles != nil {
		rolesInput := make([]string, len(*in.AllowedRoles))
		for i, role := range *in.AllowedRoles {
			rolesInput[i] = string(role)
		}
		vals["allowed_roles"] = pq.Array(rolesInput)
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
		var RoleWrapper []string
		err = rows.Scan(
			&r.ID,
			pq.Array(&r.Extensions),
			pq.Array(&RoleWrapper),
			&r.MaxSize,
			&r.UpdatedAt,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		r.AllowedRoles = make([]roles.Role, len(RoleWrapper))
		for i, role := range RoleWrapper {
			r.AllowedRoles[i], err = roles.ParseRole(role)
			if err != nil {
				return nil, fmt.Errorf("parsing role: %w", err)
			}
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
	var RoleWrapper []string
	err = q.db.QueryRowContext(ctx, query, args...).Scan(
		&r.ID,
		pq.Array(&r.Extensions),
		pq.Array(&RoleWrapper),
		&r.MaxSize,
		&r.UpdatedAt,
		&r.CreatedAt,
	)
	if err != nil {
		return MediaRulesModel{}, err
	}
	r.AllowedRoles = make([]roles.Role, len(RoleWrapper))
	for i, role := range RoleWrapper {
		r.AllowedRoles[i], err = roles.ParseRole(role)
		if err != nil {
			return MediaRulesModel{}, fmt.Errorf("parsing role: %w", err)
		}
	}

	return r, nil
}

func (q MediaRulesQ) FilterID(id string) MediaRulesQ {
	q.selector = q.selector.Where(sq.Eq{"id": id})
	q.counter = q.counter.Where(sq.Eq{"id": id})
	q.deleter = q.deleter.Where(sq.Eq{"id": id})
	q.updater = q.updater.Where(sq.Eq{"id": id})
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
