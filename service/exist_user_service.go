package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExistUserService struct {
	Qq string `json:"qq" form:"qq" binding:"required"`
}

func (e *ExistUserService) ExistUser(g *gin.Context) response.Response {

	user, err := mysql.GetUserLikeEmail(e.Qq + "@qq.com")
	if err != nil {
		util.Log().Error("exist user err: %v", err)
		return response.NewResponse(errcode.CodeUserNotExist, nil, errcode.Text(errcode.CodeUserNotExist))
	}
	if len(user) == 0 {
		return response.NewResponse(errcode.CodeUserNotExist, nil, errcode.Text(errcode.CodeUserNotExist))
	}

	reward, err := mysql.GetRewardByQq(e.Qq)

	if err != nil && err != gorm.ErrRecordNotFound {
		util.Log().Error("get reward by qq failed, err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}
	if err == gorm.ErrRecordNotFound {
		r := new(table.Reward)
		r.Count = 1
		r.Qq = e.Qq
		mysql.InsertReward(r)
	}

	if reward != nil {
		return response.NewResponse(errcode.CodeErrHasGetReward, nil, errcode.Text(errcode.CodeErrHasGetReward))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: user,
	}
}
