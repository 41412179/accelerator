package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type CountUserService struct {
}

// GetCountUser 获取用户数量
func (c *CountUserService) CountUser(ctx *gin.Context) response.Response {

	count, err := mysql.CountUser()
	if err != nil {
		util.Log().Error("get count by user id err: %v", err)
		return response.Response{
			Code: errcode.CodeDBError,
			Msg:  errcode.Text(errcode.CodeDBError),
			Data: nil,
		}
	}
	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: count,
	}
}
