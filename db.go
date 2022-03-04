package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func setupCache() (*redis.Client, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	// Mock data
	err := rdb.Set(ctx, "test-subdomain.", "127.0.0.1", 0).Err()
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