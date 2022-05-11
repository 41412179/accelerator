package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type GoodService struct {
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
	return g.setRsponse(goods)
}

// setRsponse 设置返回结果
func (g *GoodService) setRsponse(goods []*table.Good) response.Response {
	var rsponse response.Response
	rsponse.Code = errcode.CodeSuccess
	rsponse.Data = goods
	rsponse.Msg = errcode.Text(errcode.CodeSuccess)
	return rsponse
}
