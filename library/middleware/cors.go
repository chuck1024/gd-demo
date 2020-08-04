package middleware

import (
	"github.com/chuck1024/dlog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(getCorsConfig())
}

func getCorsConfig() cors.Config {
	return cors.Config{
		AllowOriginFunc: func(origin string) bool {
			dlog.Info("cors origin:%s", origin)
			//if strings.Contains(origin, "gd.com") {
			//	return true
			//}
			return false
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "User-Agent", "Cookie", "Authorization",
		},
		AllowCredentials: true,
		ExposeHeaders: []string{
			"Authorization", "Content-MD5",
		},
		MaxAge: 12 * time.Hour,
	}
}
