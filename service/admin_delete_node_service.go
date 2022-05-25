package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type AdminDeleteNodeService struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

func (a *AdminDeleteNodeService) DeleteNode(ctx *gin.Context) response.Response {
	if err := mysql.DeleteNode(a.ID); err != nil {
		util.Log().Error("delete node err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
	}
}
