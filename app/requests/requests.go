package requests

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := ctx.ShouldBind(obj); err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return false
	}

	//验证请求参数
	errs := handler(obj, ctx)

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"errors": errs})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
