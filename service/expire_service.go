package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/util"
	"time"

	"github.com/gin-gonic/gin"
)

type ExpireService struct {
	user *table.User
}

func (e *ExpireService) GetExpireTime(ctx *gin.Context) response.Response {
	user := util.GetUserByCtx(ctx)
	if user == nil {
		util.Log().Error("get user by ctx failed")
		return response.Response{
			Code: errcode.CodeCheckLogin,
			Msg:  errcode.Text(errcode.CodeCheckLogin),
		}
	}
	e.user = user

	orderService := &OrderService{}
	// 如果存在，则查询剩余时间
	remainingTime, err := orderService.GetRemainingTimeByUserId(e.user.ID)
	if err != nil {
		util.Log().Error("get remaining time by user id err: %v", err)
		return response.Response{
			Code: errcode.CodeDBError,
			Msg:  errcode.Text(errcode.CodeDBError),
			Data: nil,
		}
	}
	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: time.Now().Unix() + remainingTime,
	}
}
