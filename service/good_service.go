package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type GoodService struct {
}

type GoodInfo struct {
	// ID int64 `json:"id" form:"id" binding:"required"`
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Duration    int64     `gorm:"column:duration;type:bigint(20);NOT NULL" json:"duration"` // 用户时长
	Price       float32   `gorm:"column:price;type:bigint(20);NOT NULL" json:"price"`       // 实际单价，单位：分
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	OriginPrice float32   `gorm:"column:origin_price;type:bigint(20);NOT NULL" json:"origin_price"`
	Radio       string    `gorm:"column:radio;type:varchar(45);NOT NULL" json:"radio"` // 折扣
	Title       string    `gorm:"column:title;type:varchar(45);NOT NULL" json:"title"`
}

// GetGoods 获取商品列表
func (g *GoodService) GetGoods(c *gin.Context) response.Response {
	goods, err := mysql.GetGoods()
	if err != nil {
		util.Log().Error("get goods err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}
	if len(goods) == 0 {
		return response.NewResponse(errcode.CodeGoodNotExist, nil, errcode.Text(errcode.CodeGoodNotExist))
	}

	goodInfos := make([]*GoodInfo, 0)
	for _, good := range goods {
		goodInfo := &GoodInfo{
			Id:          good.Id,
			Duration:    good.Duration,
			Price:       float32(good.Price) / 100.0 * 6.7,
			CreatedAt:   good.CreatedAt,
			UpdatedAt:   good.UpdatedAt,
			OriginPrice: float32(good.OriginPrice) / 100.0 * 6.7,
			Radio:       good.Radio,
			Title:       fmt.Sprintf("%d", good.Duration/24/60) + "天会员",
		}
		goodInfos = append(goodInfos, goodInfo)
	}

	return g.setRsponse(goodInfos)
}

// setRsponse 设置返回结果
func (g *GoodService) setRsponse(goods []*GoodInfo) response.Response {
	var rsponse response.Response
	rsponse.Code = errcode.CodeSuccess
	rsponse.Data = goods
	rsponse.Msg = errcode.Text(errcode.CodeSuccess)
	return rsponse
}
