package session

import "github.com/go-redis/redis"

var connRedis *redis.Client

func init() {
	connRedis = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
