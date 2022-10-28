package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/config"
)

type Router struct {
	Engine *gin.Engine
}

// NewRouterは新しいRouterを初期化し構造体のポインタを返します
func NewRouter() *Router {
	e := gin.Default()
	r := &Router{
		Engine: e,
	}
	r.setMiddleware()
	return r
}

// Serveはhttpサーバーを起動します
func (r *Router) Serve() {
	r.Engine.Run(fmt.Sprintf(":%s", config.Port()))
}
