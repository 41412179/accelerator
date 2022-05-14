package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
	"time"

	"github.com/gin-gonic/gin"
)

// ProfitService 利润服务
type ProfitService struct {
	StartTime time.Time `json:"start_time" form:"start_time" binding:"required" time_format:"2006-01-02 15:04:05"`
	EndTime   time.Time `json:"end_time" form:"end_time" binding:"required" time_format:"2006-01-02 15:04:05"`
}

// ProfitServiceResponse 利润服务返回结果
type ProfitServiceResponse struct {
	TotalOrders   float64 `json:"total_orders"`
	InviterAmount float64 `json:"inviter_amount"`
	ChannelAmount float64 `json:"channel_amount"`
	ProfitAmount  float64 `json:"profit_amount"`
}

// CalcProfit 计算利润
func (p *ProfitService) CalcProfit(c *gin.Context) response.Response {
	// 获取订单金额
	orders, err := mysql.GetPaidOrdersByTime(p.StartTime, p.EndTime)
	if err != nil {
		util.Log().Error("get paid orders by time err: %+v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	r := new(ProfitServiceResponse)
	r.TotalOrders = 0
	for _, order := range orders {
		r.TotalOrders += order.PayActualPrice
		if order.ChannelId != 0 {
			r.ChannelAmount += order.PayActualPrice * 0.4
		}
	}

	// 获取邀请人利润
	inviterProfit, err := mysql.GetCommissionsByTime(p.StartTime, p.EndTime)
	if err != nil {
		util.Log().Error("get commissions by time err: %+v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	for _, profit := range inviterProfit {
		if profit.Type == table.SubCommissionType {
			r.InviterAmount += -profit.Change
		}
	}

	r.ProfitAmount = r.TotalOrders - (r.ChannelAmount + r.InviterAmount)
	return response.NewResponse(errcode.CodeSuccess, nil, errcode.Text(errcode.CodeSuccess))
}
