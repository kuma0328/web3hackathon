package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func NewSession(c *gin.Context, cookieKey string, redisValue int) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("faild to create random sentense")
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	if err := connRedis.Set(newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Error Create Session:" + err.Error())
	}
	c.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
}

func GetSession(c *gin.Context, cookieKey string) interface{} {
	redisKey, _ := c.Cookie(cookieKey)
	redisValue, err := connRedis.Get(redisKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKey not found")
		return nil
	case err != nil:
		fmt.Println("Error Get Session:" + err.Error())
		return nil
	}
	return redisValue
}

func DeleteSession(c *gin.Context, cookieKey string) {
	redisId, _ := c.Cookie(cookieKey)
	connRedis.Del(redisId)
	c.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)
}
