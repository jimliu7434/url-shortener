package model

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func Dial(addr string, psw string, db int) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: psw,
		DB:       db,
	})

	err = rdb.Ping(context.Background()).Err()
	return
}
