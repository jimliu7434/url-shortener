package model

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Dial(addr string, psw string, db int) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: psw,
		DB:       0,
	})

	err = rdb.Ping(context.Background()).Err()
	return
}
