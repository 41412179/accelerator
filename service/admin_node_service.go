package service

import (
	"accelerator/entity/db"
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type AdminNodeService struct {
}

func (a *AdminNodeService) GetNodes(ctx *gin.Context) response.Response {
	var nodes []*table.Node
	err := db.DB.Find(&nodes).Error
	if err != nil {
		util.Log().Error("get admin nodes err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: nodes,
	}
}
