package main

import (
	"cmd/internal/env"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type BikeInfo struct {
	Model string `redis:"model"`
	Brand string `redis:"brand"`
	Type  string `redis:"type"`
	Price int    `redis:"price"`
}

func main() {
	hashFields1 := []string{
		"model", "Deimos",
		"brand", "Ergonom",
		"type", "Enduro bikes",
		"price", "4972",
	}

	hashFields2 := []string{
		"model", "model",
		"brand", "brand",
		"type", "type",
		"price", "9999",
	}

	// redis
	opt, err := redis.ParseURL(env.Must("REDIS_OPT"))
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	defer client.Close()

	ctx := context.Background()

	err = client.HSet(ctx, "bike", hashFields1).Err()
	if err != nil {
		fmt.Println("Error setting Hkey:", err)
	}

	err = client.HSet(ctx, "bike", hashFields2).Err()
	if err != nil {
		fmt.Println("Error setting Hkey:", err)
	}

	var got_bike BikeInfo
	err = client.HGetAll(ctx, "bike:100").Scan(&got_bike)
	if err == redis.Nil {
		fmt.Println("NO DATA: NIL")
	} else if err != nil {
		fmt.Println("Error getting Hkey:", err)
	}

	fmt.Printf("Model: %v, Brand: %v, Type: %v, Price: %v$\n",
		got_bike.Model, got_bike.Brand, got_bike.Type, got_bike.Price)

}
