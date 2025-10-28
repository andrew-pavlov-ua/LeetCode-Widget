package redis

import (
	"cmd/internal/env"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func SetUpRDB() *redis.Client {
	opt, err := redis.ParseURL(env.Must("REDIS_OPT"))
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	// Ping to test the connection
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Could not connect to Redis:", err)
	} else {
		fmt.Println("Redis connection successful:", pong)
	}

	return client
}
