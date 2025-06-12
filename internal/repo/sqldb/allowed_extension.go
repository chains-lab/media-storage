package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
)

const allowedExtensionsTable = "allowed_extensions"

type AllowedExtensionModel struct {
	Resource  string    `db:"resource"`
	Category  string    `db:"category"`
	Extension string    `db:"extension"`
	MaxSize   int64     `db:"max_size"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AllowedExtensionQ struct {
	db       *sql.DB
	selector sq.SelectBuilder
	inserter sq.InsertBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
	counter  sq.SelectBuilder
}

func NewAllowedExtension(db *sql.DB) AllowedExtensionQ {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return AllowedExtensionQ{
		db:       db,
		selector: builder.Select("*").From(allowedExtensionsTable),
		inserter: builder.Insert(allowedExtensionsTable),
		updater:  builder.Update(allowedExtensionsTable),
		deleter:  builder.Delete(allowedExtensionsTable),
		counter:  builder.Select("COUNT(*) AS count").From(allowedExtensionsTable),
	}
}

func (q AllowedExtensionQ) New() AllowedExtensionQ {
	return NewAllowedExtension(q.db)
}

func (q AllowedExtensionQ) Insert(ctx context.Context, input AllowedExtensionModel) error {
	values := map[string]interface{}{
		"resource":   input.Resource,
		"category":   input.Category,
		"extension":  input.Extension,
		"max_size":   input.MaxSize,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
	}

	query, args, err := q.inserter.SetMap(values).ToSql()
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

type AllowedExtensionUpdateInput struct {
	Extension string    `db:"extension"`
	MaxSize   int64     `db:"max_size"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q AllowedExtensionQ) Update(ctx context.Context, input AllowedExtensionUpdateInput) error {
	values := map[string]interface{}{
		"extension":  input.Extension,
		"max_size":   input.MaxSize,
		"updated_at": input.UpdatedAt,
	}

	query, args, err := q.updater.SetMap(values).ToSql()
	if err != nil {
		return fmt.Errorf("building update query for allowed_extensions: %w", err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q AllowedExtensionQ) Get(ctx context.Context) (AllowedExtensionModel, error) {
	query, args, err := q.selector.ToSql()
	if err != nil {
		return AllowedExtensionModel{}, fmt.Errorf("building select query for allowed_extensions: %w", err)
	}

	var model AllowedExtensionModel
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		err = tx.QueryRowContext(ctx, query, args...).Scan(&model.Resource, &model.Category, &model.Extension, &model.MaxSize, &model.CreatedAt, &model.UpdatedAt)
	} else {
		err = q.db.QueryRowContext(ctx, query, args...).Scan(&model.Resource, &model.Category, &model.Extension, &model.MaxSize, &model.CreatedAt, &model.UpdatedAt)
	}

	return model, err
}

func (q AllowedExtensionQ) Select(ctx context.Context) ([]AllowedExtensionModel, error) {
	// Строим SQL-запрос
	query, args, err := q.selector.ToSql()
	if err != nil {
		return nil, fmt.Errorf("building select query for allowed_extensions: %w", err)
	}

	// Выполняем запрос в контексте транзакции или напрямую через DB
	var rows *sql.Rows
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		rows, err = tx.QueryContext(ctx, query, args...)
	} else {
		rows, err = q.db.QueryContext(ctx, query, args...)
	}
	if err != nil {
		return nil, fmt.Errorf("querying allowed_extensions: %w", err)
	}
	defer rows.Close()

	// Сканируем результаты
	var list []AllowedExtensionModel
	for rows.Next() {
		var m AllowedExtensionModel
		if err := rows.Scan(
			&m.Resource,
			&m.Category,
			&m.Extension,
			&m.MaxSize,
			&m.CreatedAt,
			&m.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scanning allowed_extensions row: %w", err)
		}
		list = append(list, m)
	}

	// Проверяем ошибки итерации
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating allowed_extensions rows: %w", err)
	}

	return list, nil
}

func (q AllowedExtensionQ) Count(ctx context.Context) (int64, error) {
	query, args, err := q.counter.ToSql()
	if err != nil {
		return 0, fmt.Errorf("building count query for allowed_extensions: %w", err)
	}

	var count int64
	err = q.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("executing count query for allowed_extensions: %w", err)
	}

	return count, nil
}

func (q AllowedExtensionQ) Delete(ctx context.Context) error {
	query, args, err := q.deleter.ToSql()
	if err != nil {
		return fmt.Errorf("building delete query for allowed_extensions: %w", err)
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = q.db.ExecContext(ctx, query, args...)
	}

	return err
}

func (q AllowedExtensionQ) FilterByResource(resource string) AllowedExtensionQ {
	q.selector = q.selector.Where(sq.Eq{"resource": resource})
	q.deleter = q.deleter.Where(sq.Eq{"resource": resource})
	q.updater = q.updater.Where(sq.Eq{"resource": resource})
	q.counter = q.counter.Where(sq.Eq{"resource": resource})
	return q
}

func (q AllowedExtensionQ) FilterByCategory(category string) AllowedExtensionQ {
	q.selector = q.selector.Where(sq.Eq{"category": category})
	q.deleter = q.deleter.Where(sq.Eq{"category": category})
	q.updater = q.updater.Where(sq.Eq{"category": category})
	q.counter = q.counter.Where(sq.Eq{"category": category})
	return q
}

func (q AllowedExtensionQ) FilterByExtension(extension string) AllowedExtensionQ {
	q.selector = q.selector.Where(sq.Eq{"extension": extension})
	q.deleter = q.deleter.Where(sq.Eq{"extension": extension})
	q.updater = q.updater.Where(sq.Eq{"extension": extension})
	q.counter = q.counter.Where(sq.Eq{"extension": extension})
	return q
}

func (q AllowedExtensionQ) Transaction(fn func(ctx context.Context) error) error {
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
