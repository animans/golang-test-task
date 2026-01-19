package storage

import "context"

type Repository interface {
	Add(ctx context.Context, v int64) error
	ListSorted(ctx context.Context) ([]int64, error)
}
