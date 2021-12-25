package infra

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/enva/service/api/internal/interface/cache"
)

type CacheHandler struct {
	Client *redis.Client
}

type CacheStatusCmd struct {
	StatusCmd *redis.StatusCmd
}

type CacheIntCmd struct {
	IntCmd *redis.IntCmd
}

type CacheStringCmd struct {
	StringCmd *redis.StringCmd
}

func NewCacheHandler(dbNum int) cache.ICacheHandler {
	cli := redis.NewClient(
		&redis.Options{
			Addr:     config.REDIS_ADDR,
			Password: config.REDIS_PASS,
			DB:       dbNum,
		},
	)
	if _, err := cli.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	return &CacheHandler{Client: cli}
}

func (c *CacheHandler) Set(ctx context.Context, key string, value interface{}, exp time.Duration) cache.IStatusCmd {
	cmd := c.Client.Set(ctx, key, value, exp)
	return &CacheStatusCmd{StatusCmd: cmd}
}

func (c *CacheHandler) Get(ctx context.Context, key string) cache.IStringCmd {
	cmd := c.Client.Get(ctx, key)
	return &CacheStringCmd{StringCmd: cmd}
}

func (c *CacheHandler) Del(ctx context.Context, args ...string) cache.IIntCmd {
	cmd := c.Client.Del(ctx, args...)
	return &CacheIntCmd{IntCmd: cmd}
}

/*   cmd methods   */

func (s *CacheStatusCmd) Result() (string, error) {
	return s.StatusCmd.Result()
}

func (s *CacheStringCmd) Result() (string, error) {
	return s.StringCmd.Result()
}

func (s *CacheIntCmd) Result() (int64, error) {
	return s.IntCmd.Result()
}
