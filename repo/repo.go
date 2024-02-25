package repo

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	RedisSet(rc *redis.Client, ctx context.Context)
}

type Repo struct {
	Title string
	Data  map[string]interface{}
}

func (r Repo) RedisSet(rc *redis.Client, ctx context.Context) {

	err := rc.Set(ctx, r.Title, r.Data, 0).Err()
	if err != nil {
		log.Printf("Failure:: Item not set.\n%v", err)
	}
}
