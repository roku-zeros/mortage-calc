package storage

import (
	"context"

	"mortage-calc/lib/cache"

)

type Storage struct {
	db *cache.Cache
}

func NewStorage(ctx context.Context) (*Storage, error) {
	db := cache.NewCache()
	return &Storage{db: db}, nil
}
