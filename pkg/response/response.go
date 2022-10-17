package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 200
func JSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

// 操作成功
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
