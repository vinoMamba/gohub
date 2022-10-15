package bootstrap

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/gohub/routes"
)

func SetupRouter(r *gin.Engine) {

	// 注册中间件
	registerMiddleware(r)

	// 注册API路由
	routes.RegisterApiRoutes(r)

	// 注册404路由
	setupNotFoundHandle(r)
}

func registerMiddleware(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setupNotFoundHandle(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		accept := ctx.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			ctx.String(404, "404 page not found")
		} else {
			ctx.JSON(404, gin.H{
				"code":    404,
				"message": "Not Found",
			})
		}
	})
}
