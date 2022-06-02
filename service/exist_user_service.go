package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type ExistUserService struct {
	Qq string `json:"qq" form:"qq" binding:"required"`
}

func (e *ExistUserService) ExistUser(g *gin.Context) response.Response {
	user, err := mysql.GetUserLikeEmail(e.Qq + "@")
	if err != nil {
		util.Log().Error("exist user err: %v", err)
		return response.NewResponse(errcode.CodeUserNotExist, nil, errcode.Text(errcode.CodeUserNotExist))
	}
	if len(user) == 0 {
		return response.NewResponse(errcode.CodeUserNotExist, nil, errcode.Text(errcode.CodeUserNotExist))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: user,
	}
}
