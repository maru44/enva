package cache

import (
	"context"
	"time"
)

type ICacheHandler interface {
	Set(context.Context, string, interface{}, time.Duration) IStatusCmd
	Get(context.Context, string) IStringCmd
	Del(context.Context, ...string) IIntCmd
}

type IStatusCmd interface {
	Result() (string, error)
}

type IStringCmd interface {
	Result() (string, error)
}

type IIntCmd interface {
	Result() (int64, error)
}
