package persistance

import "github.com/go-redis/redis"

var connRedi *redis.Client

func init() {
	connRedi = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
