package service

import (
	"accelerator/entity/response"

	"github.com/gin-gonic/gin"
)

type NodeService struct {
}

// GetNodes 获取节点列表
func (NodeService) GetNodes(c *gin.Context) response.Response {

	return response.Response{}
}
