package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
)

func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "_cookie"
		id := persistance.GetSession(c, cookieKey)
		if id == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func checkLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "_cookie"
		id := persistance.GetSession(c, cookieKey)
		if id != nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
