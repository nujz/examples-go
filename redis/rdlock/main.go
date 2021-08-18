package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func main() {
	redis.NewScript("")

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	rdb.Get(ctx, "a")
}
