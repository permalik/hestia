package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	log.Printf("Launch Sequence:: godotenv\n")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failure:: .env load\n", err)
	}

	log.Printf("Launch Sequence:: redis\n")
	connStr := os.Getenv("REDIS_CONNSTR")
	opts, err := redis.ParseURL(connStr)
	if err != nil {
		log.Fatal("Failure:: redis connection\n", err)
	}
	rc := redis.NewClient(opts)

	// compare titles: redis against gh
	k := rc.Keys(ctx, "*")
	log.Println(k)

	// d := map[string]interface{}{"asdf": "asdfsadf"}
	// r := repo.Repo{
	// 	Title: "Test",
	// 	Data:  d,
	// }
	// repo.Service.RedisSet(r, rc, ctx)

	// val, err := r.Get(ctx, "key2").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key2", val)

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":4321"))
}
