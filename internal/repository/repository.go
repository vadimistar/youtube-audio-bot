package repository

import (
	"context"
	"io"
)

type Repository interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Put(ctx context.Context, key string, file io.Reader) error
}
