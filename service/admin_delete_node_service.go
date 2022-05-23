package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type AdminDeleteNodeService struct {
	NodeID int64 `json:"node_id" form:"node_id" binding:"required"`
}

func (a *AdminDeleteNodeService) DeleteNode(ctx *gin.Context) response.Response {
	if err := mysql.DeleteNode(a.NodeID); err != nil {
		util.Log().Error("delete node err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
	}
}
