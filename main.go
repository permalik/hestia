package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func RedisClient() *redis.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed:: Load .env\n", err)
	}

	url := os.Getenv("REDIS_CONN")
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("Failed:: Connect to Redis\n", err)
	}
	return redis.NewClient(opts)
}

var ctx = context.Background()

func main() {

	r := RedisClient()

	err := r.Set(ctx, "key2", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := r.Get(ctx, "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key2", val)

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":4321"))
}
