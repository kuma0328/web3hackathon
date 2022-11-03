package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthはサーバーのヘルスチェックをするハンドラーです
func (r Router) InitHealthRouter() {
	r.Engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "good"})
	})
}
