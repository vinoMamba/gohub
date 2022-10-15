package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/vinoMamba/gohub/app/http/controllers/api/v1"
	"github.com/vinoMamba/gohub/app/models/user"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := PhoneExistRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(429, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
