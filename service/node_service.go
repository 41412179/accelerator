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
	user *table.User
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

	n.AesEncrypt(c, nodes)

	// 返回结果
	return n.setRsponse(nodes)
}

// AesEncrypt 加密节点信息
func (n *NodeService) AesEncrypt(c *gin.Context, nodes []*table.Node) {
	user := util.GetUserByCtx(c)
	vip := n.checkVIP(user, nodes)
	if !vip {
		return
	}

	var err error
	for _, node := range nodes {
		node.Link, err = util.AesEncrypt(node.Link)
		if err != nil {
			util.Log().Error("aes encrypt err: %v", err)
		}
	}
}

func (n *NodeService) checkVIP(user *table.User, nodes []*table.Node) bool {
	// 如果未登录用户
	if user == nil {
		for _, node := range nodes {
			node.Link = ""
		}
		return false
	}

	// 如果登陆用户
	o := NewOrderService()
	remainingTime, err := o.GetRemainingTimeByUserId(user.ID)
	if err != nil {
		util.Log().Error("get remaining time by user id failed, err: %v", err)
		return false
	}
	if remainingTime <= 0 {
		for _, node := range nodes {
			node.Link = ""
		}
		return false
	}

	return true

}

// setRsponse 设置返回结果
func (n *NodeService) setRsponse(nodes []*table.Node) response.Response {
	var rsponse response.Response
	rsponse.Code = errcode.CodeSuccess
	rsponse.Data = nodes
	rsponse.Msg = errcode.Text(errcode.CodeSuccess)
	return rsponse
}
