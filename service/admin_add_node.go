package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

type AdminAddNodeService struct {
	Name string `json:"name" form:"name" binding:"required"`
	Link string `json:"link" form:"link" binding:"required"`
	Tag  string `json:"tag" form:"tag" binding:"required"`
}

// AddNode
func (a *AdminAddNodeService) AddNode(c *gin.Context) response.Response {
	node := &table.Node{
		Name: a.Name,
		Link: a.Link,
		Tag:  a.Tag,
	}
	if err := mysql.AddNode(node); err != nil {
		util.Log().Error("add node err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
	}
}
