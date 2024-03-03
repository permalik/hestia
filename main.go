package main

import (
	"context"
	"github.com/google/go-github/v59/github"
	"github.com/joho/godotenv"
	"github.com/permalik/github_integration/lg"
	"github.com/permalik/github_integration/repo"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

func main() {

	// TODO: make it also update the repos per field
	lg.Launch("godotenv", nil)
	err := godotenv.Load()
	if err != nil {
		lg.Fail(".env load", "kill", err)
	}

	lg.Launch("go-github", nil)
	ghPat := os.Getenv("GITHUB_PAT")
	gc := github.NewClient(nil).WithAuthToken(ghPat)

	lg.Launch("redis", nil)
	connStr := os.Getenv("REDIS_CONNSTR")
	opts, err := redis.ParseURL(connStr)
	if err != nil {
		lg.Fail("redis connection", "kill", err)
	}
	rc := redis.NewClient(opts)

	cfg := repo.Config{
		Name: "permalik",
		Org:  false,
		Ctx:  ctx,
		Gc:   gc,
	}
	allPermalik := repo.GithubAll(cfg)
	cfg.Name = "systemysterio"
	cfg.Org = true
	allSystemysterio := repo.GithubAll(cfg)
	var allGithub []repo.Repo
	if len(allPermalik) > 0 {
		allGithub = append(allGithub, allPermalik...)
	}
	if len(allSystemysterio) > 0 {
		allGithub = append(allGithub, allSystemysterio...)
	}

	cfg.Rc = rc
	allRedis := repo.RedisAll(cfg)
	for _, v := range allRedis {
		err = repo.RedisRemoveOne(v, cfg)
		if err != nil {
			lg.Fail("RedisRemoveOne", "live", err)
		}
	}
	for _, v := range allGithub {
		err = repo.RedisAddOne(v, cfg)
		if err != nil {
			lg.Fail("RedisAddOne", "live", err)
		}
	}
	lg.Info("task complete", true, nil)
}
