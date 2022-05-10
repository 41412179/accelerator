package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

// NodeService 节点服务
type NodeService struct {
}

// GetNodes 获取节点列表
func (n *NodeService) GetNodes(c *gin.Context) response.Response {
	nodes, err := mysql.GetNodes()

	// 如果异常
	if err != nil {
		util.Log().Error("get nodes err: %v", err)
		return response.NewResponse(errcode.CodeDBError, nil, errcode.Text(errcode.CodeDBError))
	}
	if len(nodes) == 0 {
		return response.NewResponse(errcode.CodeNodeNotExist, nil, errcode.Text(errcode.CodeNodeNotExist))
	}

	n.AesEncrypt(nodes)

	// 返回结果
	return n.setRsponse(nodes)
}

// AesEncrypt 加密节点信息
func (n *NodeService) AesEncrypt(nodes []*table.Node) {
	var err error
	for _, node := range nodes {
		node.Link, err = util.AesEncrypt(node.Link)
		if err != nil {
			util.Log().Error("aes encrypt err: %v", err)
		}
	}
}

// setRsponse 设置返回结果
func (n *NodeService) setRsponse(nodes []*table.Node) response.Response {
	var rsponse response.Response
	rsponse.Code = errcode.CodeSuccess
	rsponse.Data = nodes
	rsponse.Msg = errcode.Text(errcode.CodeSuccess)
	return rsponse
}
