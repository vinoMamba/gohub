package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/vinoMamba/gohub/app/http/controllers/api/v1"
	"github.com/vinoMamba/gohub/app/models/user"
	"github.com/vinoMamba/gohub/app/requests"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(429, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	errs := requests.ValidateSignupPhoneExist(&request, ctx)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"errors": errs})
		return
	}
	ctx.JSON(200, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
