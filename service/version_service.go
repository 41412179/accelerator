package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type VersionService struct {
}

func (v *VersionService) GetVersion(ctx *gin.Context) response.Response {
	version, err := mysql.GetVersion()
	if err != nil {
		util.Log().Error("get version err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: version,
	}
}
