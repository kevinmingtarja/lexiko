package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

func setupDatabase() (*redis.Client, error) {
	ctx := context.Background()

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)

	// Mock data
	err = rdb.Set(ctx, "www.", "cname.vercel-dns.com.", 0).Err()
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, "@", "76.76.21.21", 0).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func (h *handler) getAddress(key string) (string, bool) {
	ctx := context.Background()

	val, err := h.db.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}

	return val, true
}