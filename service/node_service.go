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

// NodeService 节点服务
type NodeService struct {
}

// GetNodes 获取节点列表
func (n *NodeService) GetNodes(c *gin.Context) response.Response {
	nodes, err := mysql.GetNodes()
	// 判断是否不存在
	if err == gorm.ErrRecordNotFound {
		util.Log().Error("get nodes err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	// 如果异常
	if err != nil {
		util.Log().Error("get nodes err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}

	// 返回结果
	return n.setRsponse(nodes)
}

func (n *NodeService) setRsponse(nodes []*table.Node) response.Response {
	var rsponse response.Response
	rsponse.Code = errcode.CodeSuccess
	rsponse.Data = nodes
	rsponse.Msg = errcode.Text(errcode.CodeSuccess)
	return rsponse
}
