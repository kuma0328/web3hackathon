package persistance

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

func NewSession(c *gin.Context, cookieKey string, redisValue int) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("faild to create random sentense")
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	if err := connRedi.Set(newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Error Create Session:" + err.Error())
	}
	c.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
}

func GetSession(c *gin.Context, cookieKey string) interface{} {
	redisKey, _ := c.Cookie(cookieKey)
	redisValue, err := connRedi.Get(redisKey).Result()
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
	connRedi.Del(redisId)
	c.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)
}

// 暗号(Hash)化
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
