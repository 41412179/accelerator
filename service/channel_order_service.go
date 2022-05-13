package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"
	"time"

	"github.com/gin-gonic/gin"
)

type ChannelOrderService struct {
	ChannelId int64     `json:"channel_id" form:"channel_id" binding:"required"`
	StartTime time.Time `json:"start_time" form:"start_time" binding:"required" time_format:"2006-01-02 15:04:05"`
	EndTime   time.Time `json:"end_time" form:"end_time" binding:"required" time_format:"2006-01-02 15:04:05"`
}

type ChannelOrderServiceResponse struct {
	Number int64   `json:"number"`
	Amount float64 `json:"amount"`
}

// GetOrdersByChannelID 获取渠道佣金
func (c *ChannelOrderService) GetOrdersByChannelID(g *gin.Context) response.Response {
	orders, err := mysql.GetOrdersByChannelID(c.ChannelId, c.StartTime, c.EndTime)
	if err != nil {
		util.Log().Error("get orders by channel id err: %+v", err)
		return errcode.NewErr(errcode.CodeDBError, nil)
	}

	data := new(ChannelOrderServiceResponse)
	for _, order := range orders {
		data.Number += 1
		data.Amount += order.PayActualPrice
	}
	return response.NewResponse(errcode.CodeSuccess, data, errcode.Text(errcode.CodeSuccess))
}
