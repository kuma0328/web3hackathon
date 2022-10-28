// TODO このファイルはmiddlewareを追加するたび肥大化するので
// 本来はディレクトリを別にして分割するべき

package router

import (
	"time"

	"github.com/gin-contrib/cors"
)

// setMiddlewareはmiddlewareを設定します
func (r *Router) setMiddleware() {
	r.cors()
}

// CorsはCORSの設定を用意します
func (r *Router) cors() {
	r.Engine.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
}
