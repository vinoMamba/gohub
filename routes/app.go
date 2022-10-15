package routes

import "github.com/gin-gonic/gin"

func RegisterApiRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}
}
