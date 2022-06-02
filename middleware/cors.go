package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境需要配置跨域域名，否则403
		config.AllowOrigins = []string{"https://api.ainet.vip", "https://ainet.site", "https://www.ainet.site", "https://qaqdog.com", "http://admin.ainet.vip"}
	} else {
		config.AllowOrigins = []string{"http://apitest.ainet.site", "http://ainet.site", "http://www.ainet.site", "http://test.ainet.site:3000"}
	}
	config.AllowCredentials = true
	return cors.New(config)
}
