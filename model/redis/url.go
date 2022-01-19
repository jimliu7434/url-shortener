package model

import (
	"context"
	"fmt"
	"time"
	"url-shortener/config"
)

type TURL struct{}

var URL TURL

func (r *TURL) namespacePrefix(key string) string {
	return fmt.Sprintf("%s::%s", "URL", key)
}

func (r *TURL) SetNX(key string, value string) error {
	expireMins := config.Root.Application.DefaultExpireDays * 1440
	key = r.namespacePrefix(key)
	return rdb.SetNX(context.Background(), key, value, time.Duration(expireMins)*time.Minute).Err()
}

func (r *TURL) SetNXWithExpireTime(key string, value string, expireTime time.Time) error {
	key = r.namespacePrefix(key)
	ctx := context.Background()
	err := rdb.SetNX(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return rdb.ExpireAt(ctx, key, expireTime).Err()
}

func (r *TURL) Get(key string) (string, error) {
	key = r.namespacePrefix(key)
	return rdb.Get(context.Background(), key).Result()
}
