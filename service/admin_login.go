package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type AdminService struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (a *AdminService) AdminLogin(c *gin.Context) response.Response {
	if a.Username != "admin" || a.Password != "accelerator" {
		return response.NewResponse(errcode.CodeAdminLoginFailed, nil, errcode.Text(errcode.CodeAdminLoginFailed))
	}

	token, err := util.AesEncrypt(a.Username + ":" + a.Password)
	if err != nil {
		util.Log().Error("aes encrypt err: %v", err)
		return response.NewResponse(errcode.CodeAdminLoginFailed, nil, errcode.Text(errcode.CodeAdminLoginFailed))
	}
	return response.NewResponse(errcode.CodeSuccess, token, errcode.Text(errcode.CodeSuccess))
}
