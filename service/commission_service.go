package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

// CommissionService 佣金服务
type CommissionService struct {
	user *table.User
}

// GetCommissionByUser 获取佣金
func (c *CommissionService) GetCommissionByUser(g *gin.Context) response.Response {
	user := util.GetUserByCtx(g)
	if user == nil {
		util.Log().Error("get user by ctx failed")
		return response.Response{
			Code: errcode.CodeCheckLogin,
			Msg:  errcode.Text(errcode.CodeCheckLogin),
		}
	}
	c.user = user

	commissions, err := mysql.GetCommissionsByUser(c.user.ID)
	if err != nil {
		util.Log().Error("get commission by user failed, err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	data := c.generateData(commissions)
	return response.NewResponse(errcode.CodeSuccess, data, errcode.Text(errcode.CodeSuccess))
}

type CommissionData struct {
	CommessionAmount float64
	Number           int
}

// generateData 生成数据
func (c *CommissionService) generateData(commissions []*table.Commission) CommissionData {
	var data CommissionData
	// 佣金变化统计
	for _, commission := range commissions {
		data.CommessionAmount += commission.Change
	}

	data.Number = len(commissions)
	return data

}
