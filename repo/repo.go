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
	GithubAll(cfg Github) []Repo
	RedisAll(cfg Redis) []Repo
	RedisSet(cfg Redis)
}

type Github struct {
	Name   string
	Org    bool
	Client *github.Client
	Ctx    context.Context
}

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
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

func parseGithub(r Repo, arr []Repo, raw []*github.Repository) []Repo {

	for _, v := range raw {
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

		arr = append(arr, r)
	}
	return arr
}

func (r Repo) GithubAll(cfg Github) []Repo {

	var arr []Repo
	listOpt := github.ListOptions{Page: 1, PerPage: 25}
	if cfg.Org == true {

		opt := &github.RepositoryListByOrgOptions{Type: "public", Sort: "created", ListOptions: listOpt}
		data, _, err := cfg.Client.Repositories.ListByOrg(cfg.Ctx, cfg.Name, opt)
		if err != nil {
			lg.Fail("github: ListByOrg", err)
		}

		if len(data) <= 0 {
			lg.Info("github: no data returned from GithubAll", cfg.Name)
			return arr
		}

		arr = parseGithub(r, arr, data)
		return arr
	} else {

		opt := &github.RepositoryListByUserOptions{Type: "public", Sort: "created", ListOptions: listOpt}
		raw, _, err := cfg.Client.Repositories.ListByUser(cfg.Ctx, cfg.Name, opt)
		if err != nil {
			lg.Fail("github: ListByUser", err)
		}

		if len(raw) <= 0 {
			lg.Info("github: no data returned from GithubAll", cfg.Name)
			return arr
		}

		arr = parseGithub(r, arr, raw)
		return arr
	}
}

func (r Repo) RedisAll(cfg Redis) []Repo {

	var arr []Repo
	raw := cfg.Client.Keys(cfg.Ctx, "*")
	log.Println(raw)
	return arr
}

func (r Repo) RedisSet(cfg Redis) {

	err := cfg.Client.Set(cfg.Ctx, r.FullName, r.Data, 0).Err()
	if err != nil {
		lg.Fail("redis: Item not set", err)
	}
}
