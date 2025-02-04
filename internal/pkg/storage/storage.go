package storage

import "context"

type Storage interface {
	Store(ctx context.Context, model string, data interface{}) error
	Find(ctx context.Context, model string, filter interface{}, data interface{}) error
}
