package repo

import (
	"context"
	"github.com/google/go-github/v59/github"
	"github.com/permalik/github_integration/lg"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Service interface {
	GithubAllRepos(gc *github.Client, ctx context.Context) []Repo
	RedisSet(rc *redis.Client, ctx context.Context)
}

type Data struct {
	Name        string
	Description string
	HTMLURL     string
	Homepage    string
	Topics      []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repo struct {
	FullName string
	Data     Data
}

func (r Repo) GithubAllRepos(gc *github.Client, ctx context.Context) []Repo {

	listOpt := github.ListOptions{
		Page:    1,
		PerPage: 100,
	}

	opt := &github.RepositoryListByUserOptions{Type: "public", Sort: "created", ListOptions: listOpt}
	data, _, err := gc.Repositories.ListByUser(ctx, "permalik", opt)
	if err != nil {
		lg.Fail("github: ListByUser", err)
	}

	var rArr []Repo

	for i, v := range data {
		lg.Info("data: index", i)
		lg.Info("data: value", v.FullName)

		caTimeStamp := v.GetCreatedAt()
		caPtr := caTimeStamp.GetTime()
		createdAt := *caPtr
		uaTimeStamp := v.GetUpdatedAt()
		upPtr := uaTimeStamp.GetTime()
		updatedAt := *upPtr

		d := Data{
			Name:        v.GetName(),
			Description: v.GetDescription(),
			HTMLURL:     v.GetHTMLURL(),
			Homepage:    v.GetHomepage(),
			Topics:      v.Topics,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		r.FullName = v.GetFullName()
		r.Data = d

		rArr = append(rArr, r)
	}
	log.Println(len(rArr))
	log.Println(rArr)
	return rArr
}

func (r Repo) RedisSet(rc *redis.Client, ctx context.Context) {

	err := rc.Set(ctx, r.FullName, r.Data, 0).Err()
	if err != nil {
		lg.Fail("redis: Item not set", err)
	}
}
