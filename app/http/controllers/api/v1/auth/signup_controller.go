package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/vinoMamba/gohub/app/http/controllers/api/v1"
	"github.com/vinoMamba/gohub/app/models/user"
	"github.com/vinoMamba/gohub/app/requests"
	"github.com/vinoMamba/gohub/pkg/response"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	response.Success(ctx, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(ctx *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	response.Success(ctx, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
