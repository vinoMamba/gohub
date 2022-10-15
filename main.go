package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	//add middleware
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"message": "Not Found",
		})
	})

	r.Run(":3000")
}
