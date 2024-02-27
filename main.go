package main

import (
	"context"
	"github.com/permalik/github_integration/lg"
	"github.com/permalik/github_integration/repo"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	lg.Launch("godotenv", nil)
	err := godotenv.Load()
	if err != nil {
		lg.Fail(".env load", err)
	}

	lg.Launch("redis", nil)
	connStr := os.Getenv("REDIS_CONNSTR")
	opts, err := redis.ParseURL(connStr)
	if err != nil {
		lg.Fail("redis connection", err)
	}
	rc := redis.NewClient(opts)
	defer rc.Close()
	// lg.Launch("go-github", nil)
	// ghPat := os.Getenv("GITHUB_PAT")
	// gc := github.NewClient(nil).WithAuthToken(ghPat)

	var r repo.Repo
	// gCfg := repo.Github{
	// 	Name:   "permalik",
	// 	Org:    false,
	// 	Client: gc,
	// 	Ctx:    ctx,
	// }
	// allPermalik := repo.Service.GithubAll(r, gCfg)
	// lg.Info("all permalik", allPermalik)

	// gCfg = repo.Github{
	// 	Name:   "systemysterio",
	// 	Org:    true,
	// 	Client: gc,
	// 	Ctx:    ctx,
	// }
	// allSystemysterio := repo.Service.GithubAll(r, gCfg)
	// lg.Info("all systemysterio", allSystemysterio)

	// gCfg = repo.Github{
	// 	Name:   "azizadevelopment",
	// 	Org:    true,
	// 	Client: gc,
	// 	Ctx:    ctx,
	// }
	// allAziza := repo.Service.GithubAll(r, gCfg)
	// lg.Info("all aziza", allAziza)

	rCfg := repo.Redis{
		Client: rc,
		Ctx:    ctx,
	}
	// allRedis := repo.Service.RedisAll(r, rCfg)
	// lg.Info("all redis", allRedis)
	name := "utility"
	oneRedis := repo.Service.RedisByName(r, name, ctx, rCfg)
	lg.Info("one redis", oneRedis)

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
