package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type ShareService struct {
}

func (s *ShareService) GetShare(ctx *gin.Context) response.Response {
	share, err := mysql.GetShare()
	if err != nil {
		util.Log().Error("get share err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: share.Url,
	}
}
