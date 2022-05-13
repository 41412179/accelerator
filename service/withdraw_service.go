package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type WithdrawService struct {
	user *table.User
}

func (w *WithdrawService) WithdrawByUser(c *gin.Context) response.Response {
	user := util.GetUserByCtx(c)
	if user == nil {
		util.Log().Error("user is nil")
		return response.NewResponse(errcode.CodeCheckLogin, nil, errcode.Text(errcode.CodeCheckLogin))
	}

	// check
	// o := NewOrderService()
	change := w.GetCommissionChangeByUser(c)
	if change < 100 {
		return response.NewResponse(errcode.CodeCommissionNotEnough, nil, errcode.Text(errcode.CodeCommissionNotEnough))
	}

	// 提现
	commission := new(table.Commission)
	commission.UserId = user.ID
	commission.Change = -change
	err := mysql.InsertCommission(commission)
	if err != nil {
		util.Log().Error("insert commission failed, err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}
	return response.NewResponse(errcode.CodeSuccess, nil, errcode.Text(errcode.CodeSuccess))

}

func (w *WithdrawService) GetCommissionChangeByUser(c *gin.Context) float64 {
	commissions, err := mysql.GetCommissionsByUser(w.user.ID)
	if err != nil {
		util.Log().Error("get commission by user failed, err: %v", err)
		return 0
	}

	// data := c.generateData(commissions)
	change := 0.0

	for _, commission := range commissions {
		change += commission.Change
	}
	return change
}
