package storage

import (
	"context"
	"database/sql"
	"fmt"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Add(ctx context.Context, v int64) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO numbers(value) VALUES ($1)`, v)
	return err
}

func (r *PostgresRepo) ListSorted(ctx context.Context) ([]int64, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT value FROM numbers ORDER BY value ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []int64
	for rows.Next() {
		var v int64
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS numbers (
  id BIGSERIAL PRIMARY KEY,
  value BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_numbers_value ON numbers(value);
`)
	return err
}

func OpenPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}
