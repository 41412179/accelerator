package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
)

type CommissionService struct {
	user *table.User
}

// GetCommissionByUser 获取佣金
func (c *CommissionService) GetCommissionByUser() response.Response {
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
	number           int
}

// generateData 生成数据
func (c *CommissionService) generateData(commissions []*table.Commission) CommissionData {
	var data CommissionData
	// 佣金变化统计
	for _, commission := range commissions {
		data.CommessionAmount += commission.Change
	}

	data.number = len(commissions)
	return data

}
