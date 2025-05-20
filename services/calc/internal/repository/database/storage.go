package storage

import (
	"context"

	"github.com/roku-zeros/mortage-calc/lib/cache"
)

type Storage struct {
	db *cache.Cache
}

func NewStorage(ctx context.Context) *Storage {
	db := cache.NewCache()
	return &Storage{db: db}
}
