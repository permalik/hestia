package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/go-github/v59/github"
	"github.com/permalik/github_integration/lg"
	"github.com/redis/go-redis/v9"
	"time"
)

type Config struct {
	Name string
	Org  bool
	Ctx  context.Context
	Gc   *github.Client
	Rc   *redis.Client
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

func GithubAll(cfg Config) []Repo {

	var r Repo
	var arr []Repo
	listOpt := github.ListOptions{Page: 1, PerPage: 25}
	if cfg.Org == true {

		opt := &github.RepositoryListByOrgOptions{Type: "public", Sort: "created", ListOptions: listOpt}
		data, _, err := cfg.Gc.Repositories.ListByOrg(cfg.Ctx, cfg.Name, opt)
		if err != nil {
			lg.Fail("github: ListByOrg", "live", err)
		}

		if len(data) <= 0 {
			lg.Info("github: no data returned from GithubAll", cfg.Name)
			return arr
		}

		arr = parseGithub(r, arr, data)
		return arr
	} else {

		opt := &github.RepositoryListByUserOptions{Type: "public", Sort: "created", ListOptions: listOpt}
		raw, _, err := cfg.Gc.Repositories.ListByUser(cfg.Ctx, cfg.Name, opt)
		if err != nil {
			lg.Fail("github: ListByUser", "live", err)
		}

		if len(raw) <= 0 {
			lg.Info("github: no data returned from GithubAll", cfg.Name)
			return arr
		}
		arr = parseGithub(r, arr, raw)
		return arr
	}
}

func RedisAll(cfg Config) []string {

	res, err := cfg.Rc.Keys(cfg.Ctx, "*").Result()
	if errors.Is(err, redis.Nil) {
		lg.Info("RedisAll: redis.Nil: keys not found", err)
	} else if err != nil {
		lg.Warn("RedisAll: keys not found", err)
	}
	return res
}

func RedisAddOne(r Repo, cfg Config) error {

	data, err := json.Marshal(r.Data)
	if err != nil {
		lg.Warn("RedisSet: json.Marshal", err)
		return err
	}

	err = cfg.Rc.Set(cfg.Ctx, r.FullName, data, 0).Err()
	if err != nil {
		lg.Fail("RedisSet: Item not set", "live", err)
		return err
	}
	return nil
}

func RedisRemoveOne(fullName string, cfg Config) error {

	_, err := cfg.Rc.Del(cfg.Ctx, fullName).Result()
	if errors.Is(err, redis.Nil) {
		lg.Warn("RedisRemoveOne: name does not exist", err)
		return nil
	}
	if err != nil {
		lg.Warn("RedisRemoveOne: Rc.Del", err)
		return err
	}
	return nil
}
